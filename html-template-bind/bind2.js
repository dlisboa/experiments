/**
 * Binds event handlers for elements with g-on* attributes
 * @param {Element} element - The element to process
 * @param {Object} context - The object containing the event handler methods
 * @returns {Array} Array of attribute names to remove
 */
function bindEvents(element, context) {
  const cleanup = [];

  for (const attr of element.attributes) {
    const name = attr.name;

    // if attribute starts with 'g-on', it's an event handler
    if (name.startsWith('g-on')) {
      const eventType = name.substring(4); // Remove 'g-on' prefix
      const handler = attr.value;

      if (context && typeof context[handler] === 'function') {
        // add the event listener, binding the context
        element.addEventListener(eventType, context[handler].bind(context));

        cleanup.push(name);
      }
    }
  }

  return cleanup;
}

/**
 * Binds props for elements with g-bind attributes
 * @param {Element} element - The element to process
 * @param {Object} context - The context element to read attributes from
 * @returns {Object} Object with attributesToRemove array and propsBinding object
 */
function bindProps(element, context) {
  const cleanup = [];
  let propsBinding = null;

  for (const attr of element.attributes) {
    const name = attr.name;

    if (name === 'g-bind') {
      const bindings = attr.value.split(' ').filter(spec => spec.trim());
      const propNames = [];
      const bindingMap = new Map();

      // Parse each binding specification
      bindings.forEach(spec => {
        if (spec.includes(':')) {
          // format "from:to"
          const [from, to] = spec.split(':');
          propNames.push(from.trim());
          bindingMap.set(from.trim(), to.trim());
        } else {
          propNames.push(spec.trim());
          bindingMap.set(spec.trim(), spec.trim());
        }
      });

      // Store for two-way binding
      propsBinding = { element, propNames, bindingMap };

      // If context is a DOM element, copy the specified attributes
      if (context && context.getAttribute) {
        propNames.forEach(sourceProp => {
          const targetProp = bindingMap.get(sourceProp);
          const propValue = context.getAttribute(sourceProp);
          if (propValue !== null) {
            element.setAttribute(targetProp, propValue);
          } else {
            element.removeAttribute(targetProp);
          }
        });
      }

      // Mark g-bind attribute for removal
      cleanup.push(name);
    }
  }

  return { attributesToRemove: cleanup, propsBinding };
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
            propsBindings.forEach(({ element, propNames, bindingMap }) => {
              if (propNames.includes(attributeName)) {
                const targetAttr = bindingMap[attributeName];
                const newValue = context.getAttribute(attributeName);
                if (newValue !== null) {
                  element.setAttribute(targetAttr, newValue);
                } else {
                  element.removeAttribute(targetAttr);
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
    let cleanup = [];

    // Bind events
    const attrs = bindEvents(element, context);
    cleanup.push(...attrs);

    // Bind props
    const propsResult = bindProps(element, context);
    cleanup.push(...propsResult.attributesToRemove);
    if (propsResult.propsBinding) {
      propsBindings.push(propsResult.propsBinding);
    }

    // Bind text content for g-bind elements
    const bindElement = bindTextContent(element, context);
    if (bindElement) {
      bindElements.push(bindElement);
    }

    // Remove the processed attributes
    cleanup.forEach(attrName => {
      element.removeAttribute(attrName);
    });
  }

  // Set up two-way binding
  setupTwoWayBinding(context, propsBindings, bindElements);

  return fragment;
}