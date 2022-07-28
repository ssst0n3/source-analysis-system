export function anchor(anchorId) {
    let anchorElement = document.getElementById(anchorId)
    if (anchorElement) {
        anchorElement.scrollIntoView()
    }
}

