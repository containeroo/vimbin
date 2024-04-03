import { EditorView } from "@codemirror/view";
import { Extension } from "@codemirror/state";
import { HighlightStyle, syntaxHighlighting } from "@codemirror/language";
import { tags as t } from "@lezer/highlight";

const base00 = "#1E1E2E", // Background (Mocha base)
  base01 = "#302D41", // Overlay0 (slightly darker for distinction)
  base02 = "#575268", // Overlay1 (for secondary elements)
  base03 = "#6E6C7E", // Overlay2 (for tertiary elements)
  base04 = "#C3BAC6", // Text (Mocha text)
  base05 = "#A6ADC8", // Subtext0 (for less prominent text)
  base06 = "#D9E0EE", // Surface0 (Mocha lighter surface)
  base07 = "#F5C2E7"; // Surface1 (Mocha accent surface)

// Syntax & Accents
const base_red = "#F28FAD", // Red (Mocha red)
  base_orange = "#F8BD96", // Orange (Mocha orange)
  base_yellow = "#FAE3B0", // Yellow (Mocha yellow)
  base_green = "#ABE9B3", // Green (Mocha green)
  base_cyan = "#96CDFB", // Cyan (Mocha cyan)
  base_blue = "#89DCEB", // Blue (Mocha blue)
  base_violet = "#DDB6F2", // Violet (Mocha violet)
  base_magenta = "#F5C2E7"; // Magenta (Mocha magenta)

const invalid = "#EC0101", // Adjusted to Mocha's emphasis color for errors
  stone = base04, // Using Text color as stone equivalent
  darkBackground = "#181825", // Darker variant of the background
  highlightBackground = "#29283C", // Highlight background (slightly lighter than base01 for visibility)
  background = base00, // Background (Mocha base)
  tooltipBackground = base01, // Using Overlay0 for tooltips
  selection = "#414052", // Selection color, slightly lighter than Overlay1
  cursor = base04; // Using Text color for cursor

/// The editor theme styles for Solarized Dark.
export const catppuccinMochaTheme = EditorView.theme(
  {
    "&": {
      color: base05,
      backgroundColor: background,
    },

    ".cm-content": {
      caretColor: cursor,
    },

    ".cm-cursor, .cm-dropCursor": { borderLeftColor: cursor },
    "&.cm-focused > .cm-scroller > .cm-selectionLayer .cm-selectionBackground, .cm-selectionBackground, .cm-content ::selection":
      { backgroundColor: selection },

    ".cm-panels": { backgroundColor: darkBackground, color: base03 },
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
      outline: `1px solid ${base06}`,
    },

    ".cm-gutters": {
      backgroundColor: darkBackground,
      color: stone,
      border: "none",
    },

    ".cm-activeLineGutter": {
      backgroundColor: highlightBackground,
    },

    ".cm-foldPlaceholder": {
      backgroundColor: "transparent",
      border: "none",
      color: "#ddd",
    },

    ".cm-tooltip": {
      border: "none",
      backgroundColor: tooltipBackground,
    },
    ".cm-tooltip .cm-tooltip-arrow:before": {
      borderTopColor: "transparent",
      borderBottomColor: "transparent",
    },
    ".cm-tooltip .cm-tooltip-arrow:after": {
      borderTopColor: tooltipBackground,
      borderBottomColor: tooltipBackground,
    },
    ".cm-tooltip-autocomplete": {
      "& > ul > li[aria-selected]": {
        backgroundColor: highlightBackground,
        color: base03,
      },
    },
  },
  { dark: true },
);

/// The highlighting style for code in the Solarized Dark theme.
export const catppuccinMochaHighlightStyle = HighlightStyle.define([
  { tag: t.keyword, color: base_green },
  {
    tag: [t.name, t.deleted, t.character, t.propertyName, t.macroName],
    color: base_cyan,
  },
  { tag: [t.variableName], color: base05 },
  { tag: [t.function(t.variableName)], color: base_blue },
  { tag: [t.labelName], color: base_magenta },
  {
    tag: [t.color, t.constant(t.name), t.standard(t.name)],
    color: base_yellow,
  },
  { tag: [t.definition(t.name), t.separator], color: base_cyan },
  { tag: [t.brace], color: base_magenta },
  {
    tag: [t.annotation],
    color: invalid,
  },
  {
    tag: [t.number, t.changed, t.annotation, t.modifier, t.self, t.namespace],
    color: base_magenta,
  },
  {
    tag: [t.typeName, t.className],
    color: base_orange,
  },
  {
    tag: [t.operator, t.operatorKeyword],
    color: base_violet,
  },
  {
    tag: [t.tagName],
    color: base_blue,
  },
  {
    tag: [t.squareBracket],
    color: base_red,
  },
  {
    tag: [t.angleBracket],
    color: base02,
  },
  {
    tag: [t.attributeName],
    color: base05,
  },
  {
    tag: [t.regexp],
    color: invalid,
  },
  {
    tag: [t.quote],
    color: base_green,
  },
  { tag: [t.string], color: base_yellow },
  {
    tag: t.link,
    color: base_cyan,
    textDecoration: "underline",
    textUnderlinePosition: "under",
  },
  {
    tag: [t.url, t.escape, t.special(t.string)],
    color: base_yellow,
  },
  { tag: [t.meta], color: base_red },
  { tag: [t.comment], color: base02, fontStyle: "italic" },
  { tag: t.strong, fontWeight: "bold", color: base06 },
  { tag: t.emphasis, fontStyle: "italic", color: base_green },
  { tag: t.strikethrough, textDecoration: "line-through" },
  { tag: t.heading, fontWeight: "bold", color: base_yellow },
  { tag: t.heading1, fontWeight: "bold", color: base07 },
  {
    tag: [t.heading2, t.heading3, t.heading4],
    fontWeight: "bold",
    color: base06,
  },
  {
    tag: [t.heading5, t.heading6],
    color: base06,
  },
  { tag: [t.atom, t.bool, t.special(t.variableName)], color: base_magenta },
  {
    tag: [t.processingInstruction, t.inserted, t.contentSeparator],
    color: base_red,
  },
  {
    tag: [t.contentSeparator],
    color: base_yellow,
  },
  { tag: t.invalid, color: base02, borderBottom: `1px dotted ${base_red}` },
]);

/// Extension to enable the Solarized Dark theme (both the editor theme and
/// the highlight style).
export const catppuccinMocha: Extension = [
  catppuccinMochaTheme,
  syntaxHighlighting(catppuccinMochaHighlightStyle),
];
