/**
 * Processes <bind> elements and applies bindings to their children
 * @param {Element} element - The element to process
 * @param {Object} context - The context element to read attributes from
 * @returns {Object} Object with bindings information and elements to remove
 */
const TAG_NAME = 'bind'
function processBindElement(element, context) {
  if (element.tagName.toLowerCase() !== TAG_NAME) {
    return { propsBindings: [], bindElements: [], elementsToRemove: [] };
  }

  const propsBindings = [];
  const bindElements = [];
  const elementsToRemove = [element];

  // Check for special 'self' attribute for text content binding
  const selfAttr = element.getAttribute('self');
  if (selfAttr !== null) {
    // This is a text content binding
    const propName = selfAttr || 'self'; // If self="", use "self" as default

    if (context && context.getAttribute) {
      // Set initial text content from context attribute
      const propValue = context.getAttribute(propName);
      element.textContent = propValue || '';

      // Store for two-way binding
      bindElements.push({ element, propName });
    }

    return { propsBindings, bindElements, elementsToRemove: [] }; // Don't remove self-binding elements
  }

  // Process attribute bindings for child elements
  const children = Array.from(element.children);

  children.forEach(child => {
    const bindingMap = {};
    const propNames = [];

    // Process all attributes on the <bind> element
    for (let i = 0; i < element.attributes.length; i++) {
      const attr = element.attributes[i];
      const attrName = attr.name;
      const attrValue = attr.value;

      // Handle event bindings (onclick, onchange, etc.)
      if (attrName.startsWith('on')) {
        const eventType = attrName.substring(2); // Remove 'on' prefix
        const handlerName = attrValue;

        // Check if the context object has a method with this name
        if (context && typeof context[handlerName] === 'function') {
          // Add the event listener, binding the context as 'this'
          child.addEventListener(eventType, context[handlerName].bind(context));
        }
      } else {
        // Handle attribute bindings
        let sourceProp, targetProp;

        if (attrValue === '' || attrValue === attrName) {
          // Empty value or same as attribute name: use attribute name as both source and target
          sourceProp = attrName;
          targetProp = attrName;
        } else {
          // Explicit value: use value as source, attribute name as target
          sourceProp = attrValue;
          targetProp = attrName;
        }

        propNames.push(sourceProp);
        bindingMap[sourceProp] = targetProp;

        // If context is a DOM element, copy the specified attributes
        if (context && context.getAttribute) {
          const propValue = context.getAttribute(sourceProp);
          if (propValue !== null) {
            child.setAttribute(targetProp, propValue);
          } else {
            child.removeAttribute(targetProp);
          }
        }
      }
    }

    // Store bindings for two-way binding if we have any
    if (propNames.length > 0) {
      propsBindings.push({ element: child, propNames, bindingMap });
    }
  });

  // Replace <bind> element with its children
  const parent = element.parentNode;
  if (parent) {
    children.forEach(child => {
      parent.insertBefore(child, element);
    });
  }

  return { propsBindings, bindElements, elementsToRemove };
}

/**
 * Sets up two-way binding with MutationObserver
 * @param {Object} context - The context element to observe
 * @param {Array} propsBindings - Array of props bindings
 * @param {Array} bindElements - Array of bind elements
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

            // Update all bind elements that use this attribute
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

  // Process <bind> elements
  const bindElements = Array.from(fragment.querySelectorAll(TAG_NAME));
  const allPropsBindings = [];
  const allBindElements = [];

  bindElements.forEach(bindElement => {
    const result = processBindElement(bindElement, context);
    allPropsBindings.push(...result.propsBindings);
    allBindElements.push(...result.bindElements);

    // Remove the <bind> elements from DOM after processing
    result.elementsToRemove.forEach(elementToRemove => {
      if (elementToRemove.parentNode) {
        elementToRemove.parentNode.removeChild(elementToRemove);
      }
    });
  });

  // Set up two-way binding
  setupTwoWayBinding(context, allPropsBindings, allBindElements);

  return fragment;
}