// ssr.js
import { createRequire } from 'module';
const require = createRequire(import.meta.url);
const React = require('react');
const ReactDOMServer = require('react-dom/server');

let input = '';
process.stdin.setEncoding('utf8');
process.stdin.on('data', chunk => (input += chunk));
process.stdin.on('end', () => {
  try {
    const exports = {};
    const module = { exports };
    const fn = new Function('require', 'exports', 'module', input);
    fn(require, exports, module);
    const Component = module.exports.default;
    const html = ReactDOMServer.renderToString(React.createElement(Component));
    process.stdout.write(html);
  } catch (err) {
    console.error('SSR Error:', err);
    process.exit(1);
  }
});
 