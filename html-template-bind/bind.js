/**
 * Binds event handlers for elements with g-on* attributes
 * @param {Element} element - The element to process
 * @param {Object} context - The object containing the event handler methods
 * @returns {Array} Array of attribute names to remove
 */
function bindEvents(element, context) {
  const attributesToRemove = [];

  // Check all attributes on this element
  for (let i = 0; i < element.attributes.length; i++) {
    const attr = element.attributes[i];
    const attrName = attr.name;

    // If attribute starts with 'g-on', it's an event handler
    if (attrName.startsWith('g-on')) {
      const eventType = attrName.substring(4); // Remove 'g-on' prefix
      const handlerName = attr.value;

      // Check if the context object has a method with this name
      if (context && typeof context[handlerName] === 'function') {
        // Add the event listener, binding the context as 'this'
        element.addEventListener(eventType, context[handlerName].bind(context));

        // Mark this attribute for removal
        attributesToRemove.push(attrName);
      }
    }
  }

  return attributesToRemove;
}

/**
 * Binds props for elements with g-bind attributes
 * @param {Element} element - The element to process
 * @param {Object} context - The context element to read attributes from
 * @returns {Object} Object with attributesToRemove array and propsBinding object
 */
function bindProps(element, context) {
  const attributesToRemove = [];
  let propsBinding = null;

  // Check all attributes on this element
  for (let i = 0; i < element.attributes.length; i++) {
    const attr = element.attributes[i];
    const attrName = attr.name;

    // Handle g-bind attribute for passing down attributes from context element
    if (attrName === 'g-bind') {
      const propNames = attr.value.split(' ').filter(name => name.trim());

      // Store for two-way binding
      propsBinding = { element, propNames };

      // If context is a DOM element, copy the specified attributes
      if (context && context.getAttribute) {
        propNames.forEach(propName => {
          const propValue = context.getAttribute(propName);
          if (propValue !== null) {
            element.setAttribute(propName, propValue);
          } else {
            element.removeAttribute(propName);
          }
        });
      }

      // Mark g-bind attribute for removal
      attributesToRemove.push(attrName);
    }
  }

  return { attributesToRemove, propsBinding };
}

/**
 * Binds g-bind elements for text content binding
 * @param {Element} element - The element to process
 * @param {Object} context - The context element to read attributes from
 * @returns {Object|null} Bind element object or null
 */
function bindTextContent(element, context) {
  // Handle g-bind elements
  if (element.tagName.toLowerCase() === 'g-bind') {
    const propAttr = element.getAttribute('prop');
    if (propAttr && context && context.getAttribute) {
      // Set initial text content from context attribute
      const propValue = context.getAttribute(propAttr);
      element.textContent = propValue || '';

      // Return binding info for two-way binding
      return { element, propName: propAttr };
    }
  }

  return null;
}

/**
 * Sets up two-way binding with MutationObserver
 * @param {Object} context - The context element to observe
 * @param {Array} propsBindings - Array of props bindings
 * @param {Array} bindElements - Array of g-bind elements
 */
function setupTwoWayBinding(context, propsBindings, bindElements) {
  if ((propsBindings.length > 0 || bindElements.length > 0) && context && context.getAttribute) {
    // Check if we already have an observer to avoid duplicates
    if (!context._bindObserver) {
      const observer = new MutationObserver((mutations) => {
        mutations.forEach((mutation) => {
          if (mutation.type === 'attributes') {
            const attributeName = mutation.attributeName;

            // Update all bound elements that use this attribute
            propsBindings.forEach(({ element, propNames }) => {
              if (propNames.includes(attributeName)) {
                const newValue = context.getAttribute(attributeName);
                if (newValue !== null) {
                  element.setAttribute(attributeName, newValue);
                } else {
                  element.removeAttribute(attributeName);
                }
              }
            });

            // Update all g-bind elements that use this attribute
            bindElements.forEach(({ element, propName }) => {
              if (propName === attributeName) {
                const newValue = context.getAttribute(attributeName);
                element.textContent = newValue || '';
              }
            });
          }
        });
      });

      observer.observe(context, {
        attributes: true,
        attributeOldValue: false
      });

      // Store observer reference to avoid creating duplicates
      context._bindObserver = observer;

      // Store bindings on the context for this observer
      if (!context._bindingsRefs) {
        context._bindingsRefs = [];
      }
      if (!context._bindElementsRefs) {
        context._bindElementsRefs = [];
      }
    }

    // Add current bindings to the context's bindings list
    if (context._bindingsRefs) {
      context._bindingsRefs.push(...propsBindings);
    }
    if (context._bindElementsRefs) {
      context._bindElementsRefs.push(...bindElements);
    }
  }
}

/**
 * Binds event handlers from a template to an object's methods and passes props from context to elements
 * @param {HTMLTemplateElement} template - The template element to process
 * @param {Object} context - The object containing the event handler methods (should be a DOM element to read attributes from)
 * @returns {DocumentFragment} A document fragment with event listeners attached and props bound
 */
export function bind(template, context) {
  // Clone the template content to avoid modifying the original
  const fragment = template.content.cloneNode(true);

  // Create a TreeWalker to traverse all nodes in the fragment
  const walker = document.createTreeWalker(
    fragment,
    NodeFilter.SHOW_ELEMENT, // Only process element nodes
    null,
    false
  );

  let node;
  const propsBindings = []; // Store elements and their props for two-way binding
  const bindElements = []; // Store g-bind elements for text content binding

  // Walk through all element nodes in the fragment
  while (node = walker.nextNode()) {
    const element = node;
    let attributesToRemove = [];

    // Bind events
    let eventAttrsToRemove = bindEvents(element, context);
    attributesToRemove.push(...eventAttrsToRemove);

    // Bind props
    const propsResult = bindProps(element, context);
    attributesToRemove.push(...propsResult.attributesToRemove);
    if (propsResult.propsBinding) {
      propsBindings.push(propsResult.propsBinding);
    }

    // Bind text content for g-bind elements
    const bindElement = bindTextContent(element, context);
    if (bindElement) {
      bindElements.push(bindElement);
    }

    // Remove the processed attributes
    attributesToRemove.forEach(attrName => {
      element.removeAttribute(attrName);
    });
  }

  // Set up two-way binding
  setupTwoWayBinding(context, propsBindings, bindElements);

  return fragment;
}