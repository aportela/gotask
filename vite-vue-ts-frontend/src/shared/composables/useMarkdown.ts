import MarkdownIt from "markdown-it";
import DOMPurify from "dompurify";

import hljs from "highlight.js";
import "highlight.js/styles/github.css";

// @ts-expect-error
import TurndownService from "turndown";

const md = new MarkdownIt({
  html: false,
  linkify: true,
  typographer: true,
  breaks: true,
});

md.options.highlight = (code: string, lang: string): string => {
  if (lang && hljs.getLanguage(lang)) {
    return `<pre><code class="hljs ${lang}">${hljs.highlight(code, { language: lang }).value}</code></pre>`;
  }
  return `<pre><code class="hljs">${md.utils.escapeHtml(code)}</code></pre>`;
};

const turndown = new TurndownService();

export function useMarkdown() {
  const render = (text: string) => DOMPurify.sanitize(md.render(text));

  const toMarkdown = (html: string) => turndown.turndown(html);

  return {
    md,
    turndown,
    render,
    toMarkdown,
  };
}
