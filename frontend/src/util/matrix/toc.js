import {marked} from "../marked";

function find_headings(doc) {
    let headings = []
    let tag_names = {
        h1: 1,
        h2: 1,
        h3: 1,
        h4: 1,
        h5: 1,
        h6: 1
    }
    let title = null

    function walk(root) {
        if (root.nodeType === 1 && root.nodeName !== 'script') {
            if (Object.prototype.hasOwnProperty.call(tag_names, root.nodeName.toLowerCase())) {
                if (root.getAttribute("class") === null) {
                    if (title === null && root.nodeName.toLowerCase() === "h1") {
                        title = root
                    } else {
                        headings.push(root);
                    }
                }
            } else {
                for (let i = 0; i < root.childNodes.length; i++) {
                    walk(root.childNodes[i]);
                }
            }
        }
    }

    walk(doc);
    return {
        title: title,
        headings: headings,
    }
}

function parseHeading(markdown) {
    let rendered = marked.parse(markdown)
    let div = document.createElement('div')
    div.innerHTML = rendered
    return find_headings(div)
}

export {parseHeading}