import {Marked} from "marked";
import {markedHighlight} from "marked-highlight";
import {gfmHeadingId } from "marked-gfm-heading-id";
import hljs from 'highlight.js';

const marked = new Marked(
    markedHighlight({
        langPrefix: 'hljs language-',
        highlight(code, lang) {
            const language = hljs.getLanguage(lang) ? lang : 'plaintext';
            return hljs.highlight(code, {language}).value;
        }
    })
)
marked.use(gfmHeadingId());

export {marked}
