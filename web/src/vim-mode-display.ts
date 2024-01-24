import { ViewUpdate, ViewPlugin, PluginValue } from "@codemirror/view";
import { getCM } from "@replit/codemirror-vim";

export function vimModeDisplay() {
  return ViewPlugin.fromClass(
    class implements PluginValue {
      update(update: ViewUpdate) {
        const { view } = update;
        const cm = getCM(view);

        if (!cm) {
          return;
        }

        // Retrieve the current Vim mode
        const currentVimMode = cm.state.vim?.mode || "unknown";

        // Update the Vim mode display element
        const modeElement = document.getElementById("vim-mode-status");
        if (modeElement) {
          modeElement.textContent = currentVimMode.toUpperCase();

          // Update the class for styling
          modeElement.className = `vim-mode ${currentVimMode}`;
        }
      }

      destroy() {
        // Cleanup logic if needed
      }
    },
  );
}
