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

  // Walk through all element nodes in the fragment
  while (node = walker.nextNode()) {
    const element = node;
    const attributesToRemove = [];

    // Check all attributes on this element
    for (let i = 0; i < element.attributes.length; i++) {
      const attr = element.attributes[i];
      const attrName = attr.name;

      // If attribute starts with 'on', it's an event handler
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

      // Handle props attribute for passing down attributes from context element
      if (attrName === 'g-props') {
        const propNames = attr.value.split(' ').filter(name => name.trim());

        // Store for two-way binding
        propsBindings.push({ element, propNames });

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

        // Mark props attribute for removal
        attributesToRemove.push(attrName);
      }
    }

    // Remove the processed attributes
    attributesToRemove.forEach(attrName => {
      element.removeAttribute(attrName);
    });
  }

  // Set up two-way binding with MutationObserver
  if (propsBindings.length > 0 && context && context.getAttribute) {
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
    }

    // Add current bindings to the context's bindings list
    if (context._bindingsRefs) {
      context._bindingsRefs.push(...propsBindings);
    }
  }

  return fragment;
}
