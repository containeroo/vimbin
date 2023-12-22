import {
  autocompletion,
  closeBrackets,
  closeBracketsKeymap,
  completionKeymap,
} from "@codemirror/autocomplete";
import { defaultKeymap, history, historyKeymap } from "@codemirror/commands";
import { vim, Vim } from "@replit/codemirror-vim";
import {
  bracketMatching,
  defaultHighlightStyle,
  foldGutter,
  foldKeymap,
  indentOnInput,
  syntaxHighlighting,
} from "@codemirror/language";
import { lintKeymap } from "@codemirror/lint";
import { highlightSelectionMatches, searchKeymap } from "@codemirror/search";
import { EditorState } from "@codemirror/state";
import {
  crosshairCursor,
  drawSelection,
  dropCursor,
  EditorView,
  highlightActiveLine,
  highlightActiveLineGutter,
  highlightSpecialChars,
  keymap,
  lineNumbers,
  rectangularSelection,
} from "@codemirror/view";

import { Extension } from "@codemirror/state";

export const lineNumbersRelative: Extension = [formatNumber()];

function formatNumber() {
  return lineNumbers({
    formatNumber: (lineNo, state) => {
      if (lineNo > state.doc.lines) {
        return "0";
      }
      const cursorLine = state.doc.lineAt(
        state.selection.asSingle().ranges[0].to,
      ).number;
      if (lineNo === cursorLine) {
        return "0";
      } else {
        return Math.abs(cursorLine - lineNo).toString();
      }
    },
  });
}

Vim.defineEx("x", "x", function () {
  console.log("write");
});

Vim.defineEx("write", "w", function () {
  console.log("write");
});

export async function initializeEditor(initialContent: string) {
  const state = EditorState.create({
    doc: initialContent,
    extensions: [
      vim(),
      lineNumbersRelative,
      highlightActiveLineGutter(),
      highlightSpecialChars(),
      history(),
      foldGutter(),
      drawSelection(),
      dropCursor(),
      EditorState.allowMultipleSelections.of(true),
      indentOnInput(),
      syntaxHighlighting(defaultHighlightStyle, { fallback: true }),
      bracketMatching(),
      closeBrackets(),
      autocompletion(),
      rectangularSelection(),
      crosshairCursor(),
      highlightActiveLine(),
      highlightSelectionMatches(),
      keymap.of([
        ...closeBracketsKeymap,
        ...defaultKeymap,
        ...searchKeymap,
        ...historyKeymap,
        ...foldKeymap,
        ...completionKeymap,
        ...lintKeymap,
      ]),
    ],
  });

  const targetElement = document.querySelector("#editor")!;
  const editor = new EditorView({
    parent: targetElement,
    state: state,
  });

  editor.focus();
}
