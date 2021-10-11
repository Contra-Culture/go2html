# go2html

Go2HTML is a library for making and organizing HTML templates.

## Glossary

**Document (page) *[not implemented yet]*** is a highest-level `Template` that is responsible for whole page generation. The purpose of a `Document` is to provide a context for valid complete page generation.

**Template** - is a function for valid HTML page fragment generation. Template consists of `Nodes` and uses precompillation for HTML generation performance improvement (instead of tree traversalit uses fragments slice of strings and `injections` to perform minimal work for template population).

**Injection** - is a mark within a `Template` where content should be placed during HTML code generation by the the `Template` population.

**Node** - is a context-dependent building block. Do not confuse DOM API Nodes with *go2html* nodes! In this document we use term `Node` only for *go2html*'s `Nodes`. `Template` is a tree built of `Nodes`. There is sevaral types of `Nodes`:

- *`element`* - adds an HTML element `Node` into `Template` or parent HTML element `Node` context.
- *`template`* - adds a `Template` embedding `Node` into `Template` or parent HTML element `Node` context.
- *`comment`* - adds a HTML comment into `Template` or parent HTML element `Node` context.
- *`doctype`* - adds a `DOCTYPE` definition into `Template`. Supports only HTML5 definition. [no validations implemented yet].
- *`text-injection`* - adds safe text injection `Node` into `Template` or parent HTML element `Node` context. The injection will be replaced with exact text, provided during `Template` populating. HTML will be escaped.
- *`unsafe-text-injection`* - adds **UNSAFE** text injection `Node` into `Template` or parent HTML element `Node` context. The injection will be replaced with exact text, provided during `Template` populating. HTML **WILL NOT** be escaped. Try to avoid unsafe text injections. Use it only for text that comes from the sources you really trust.
- *`text`* - add safe text node into `Template` or parent HTML element `Node` with exact text. HTML will be escaped.
- *`unsafe-text`* - add **UNSAFE** text `Node` into `Template` or parent HTML element `Node` with exact text. HTML **WILL NOT** be escaped. Try to avoid unsafe text nodes. Use it only for text that comes from the sources you really trust.
- *`repeat`* - adds repeatable `Template` embedding  `Node` into `Template` or parent HTML element `Node`. For example, if you want to show a list of articles - use `repeat` together with a `Template` that is responsible for single article preview generation.
- *`styling` [not implemented yet]* - allows to generate CSS from HTML.
- *`variant` / `condition` [not implemented yet]* - allows to define several candidates for injection into resulting HTML, depending on `Template`'s populating data/parameters.

## Project Purpose

go2html is a tool for HTML templates generation. It provides all the abilities to generate valid and safe HTML code, reuse and organize HTML templates, extend them with features, like CSS from HTML generation that helps to avoid dead CSS code. *go2html* also allows you to keep your design system's artefacts in code.

## Usage

[TODO] See tests.
## Best Practices
