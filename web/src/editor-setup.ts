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
import { vimModeDisplay } from "./vim-mode-display";
import { Compartment } from "@codemirror/state";
import { catppuccinFrappe } from "./catppuccine-frappe";

const themeConfig = new Compartment();

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

function createEditor() {
  const targetElement = document.querySelector("#editor")!;
  return new EditorView({
    parent: targetElement,
  });
}

export const editor = createEditor();

export async function initializeEditor(initialContent: string) {
  const editorView = editor;

  const state = EditorState.create({
    doc: initialContent,
    extensions: [
      EditorState.allowMultipleSelections.of(true),
      catppuccinFrappe,
      autocompletion(),
      bracketMatching(),
      closeBrackets(),
      crosshairCursor(),
      drawSelection(),
      dropCursor(),
      foldGutter(),
      highlightActiveLine(),
      highlightActiveLineGutter(),
      highlightSelectionMatches(),
      highlightSpecialChars(),
      history(),
      indentOnInput(),
      lineNumbersRelative,
      rectangularSelection(),
      syntaxHighlighting(defaultHighlightStyle, { fallback: true }),
      vim(),
      vimModeDisplay(),
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

  editorView.setState(state);

  Vim.exitInsertMode((editorView as any).cm);
  editorView.focus();
}
