import { EditorView } from "@codemirror/view";
import { Extension } from "@codemirror/state";
import { HighlightStyle, syntaxHighlighting } from "@codemirror/language";
import { tags as t } from "@lezer/highlight";

// Colors adapted from your CSS variables to TypeScript format
const base = "rgb(48, 52, 70)", // --ctp-base
  overlay1 = "rgb(131, 138, 164)", // --ctp-overlay1
  overlay0 = "rgb(115, 120, 145)", // --ctp-overlay0
  surface0 = "rgb(65, 69, 89)", // --ctp-surface0
  text = "rgb(198, 206, 239)", // --ctp-text
  subtext0 = "rgb(165, 172, 201)", // --ctp-subtext0
  red = "rgb(231, 130, 132)", // --ctp-red
  orange = "rgb(239, 159, 118)", // --ctp-peach to match your usage of orange
  yellow = "rgb(229, 200, 144)", // --ctp-yellow
  green = "rgb(166, 209, 137)", // --ctp-green
  cyan = "rgb(153, 209, 219)", // --ctp-sky to match your usage of cyan
  blue = "rgb(140, 170, 238)", // --ctp-blue
  violet = "rgb(202, 158, 230)", // --ctp-mauve to match your usage of violet
  magenta = "rgb(244, 184, 228)"; // --ctp-pink to match your usage of magenta

// Syntax & Accents
const invalid = "#EC0101", // Adjusted for visibility
  highlightBackground = "rgb(81, 86, 108)", // --ctp-surface1
  selection = "rgb(98, 103, 126)", // --ctp-surface2
  cursor = text; // Use text color for cursor for consistency

/// The editor theme styles.
export const catppuccinFrappeTheme = EditorView.theme(
  {
    "&": {
      color: text,
      backgroundColor: base,
    },

    ".cm-content": {
      caretColor: cursor,
    },

    ".cm-cursor, .cm-dropCursor": { borderLeftColor: cursor },
    "&.cm-focused .cm-selectionBackground, .cm-selectionBackground, .cm-content ::selection":
      {
        backgroundColor: selection,
      },

    ".cm-panels": { backgroundColor: overlay1, color: subtext0 },
    ".cm-panels.cm-panels-top": { borderBottom: "2px solid black" },
    ".cm-panels.cm-panels-bottom": { borderTop: "2px solid black" },

    ".cm-searchMatch": {
      backgroundColor: "#72a1ff59",
      outline: "1px solid #457dff",
    },
    ".cm-searchMatch.cm-searchMatch-selected": {
      backgroundColor: "#6199ff2f",
    },

    ".cm-activeLine": { backgroundColor: highlightBackground },
    ".cm-selectionMatch": { backgroundColor: "#aafe661a" },

    "&.cm-focused .cm-matchingBracket, &.cm-focused .cm-nonmatchingBracket": {
      outline: `1px solid ${yellow}`,
    },

    ".cm-gutters": {
      backgroundColor: overlay0,
      color: subtext0,
      border: "none",
    },

    ".cm-activeLineGutter": {
      backgroundColor: highlightBackground,
    },

    ".cm-tooltip": {
      border: "none",
      backgroundColor: surface0,
    },
    ".cm-tooltip-autocomplete": {
      "& > ul > li[aria-selected]": {
        backgroundColor: highlightBackground,
        color: text,
      },
    },
  },
  { dark: true },
);

/// The highlighting style for code.
export const catppuccinFrappeHighlightStyle = HighlightStyle.define([
  { tag: t.keyword, color: violet },
  {
    tag: [t.name, t.deleted, t.character, t.propertyName, t.macroName],
    color: blue,
  },
  { tag: [t.variableName], color: text },
  { tag: [t.function(t.variableName)], color: orange },
  { tag: [t.labelName], color: magenta },
  { tag: [t.color, t.constant(t.name), t.standard(t.name)], color: yellow },
  { tag: [t.definition(t.name), t.separator], color: cyan },
  { tag: [t.brace], color: magenta },
  { tag: [t.annotation], color: invalid },
  {
    tag: [t.number, t.changed, t.annotation, t.modifier, t.self, t.namespace],
    color: orange,
  },
  { tag: [t.typeName, t.className], color: green },
  { tag: [t.operator, t.operatorKeyword], color: red },
  { tag: [t.tagName], color: orange },
  { tag: [t.squareBracket], color: red },
  { tag: [t.angleBracket], color: overlay1 },
  { tag: [t.attributeName], color: green },
  { tag: [t.regexp], color: invalid },
  { tag: [t.quote], color: yellow },
  { tag: [t.string], color: green },
  { tag: t.link, color: cyan, textDecoration: "underline" },
  { tag: [t.meta], color: red },
  { tag: [t.comment], color: overlay1, fontStyle: "italic" },
  { tag: t.invalid, color: invalid, borderBottom: `1px dotted ${red}` },
]);

/// Extension to enable the theme (both the editor theme and the highlight style).
export const catppuccinFrappe: Extension = [
  catppuccinFrappeTheme,
  syntaxHighlighting(catppuccinFrappeHighlightStyle),
];
