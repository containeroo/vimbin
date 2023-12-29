import { ViewUpdate, ViewPlugin, PluginValue } from "@codemirror/view";
import { Vim, getCM } from "@replit/codemirror-vim";

export function vimModeDisplay() {
  return ViewPlugin.fromClass(
    class implements PluginValue {
      update(update: ViewUpdate) {
        const { view } = update;
        const cm = getCM(view);

        if (!cm) {
          return;
        }

        // Log the current Vim mode
        console.log(cm.getMode());
        // logCurrentVimMode(cm);
      }

      destroy() {
        // Cleanup logic if needed
      }
    },
    {
      // Add other configuration options if necessary
    },
  );
}

function logCurrentVimMode(cm: any) {
  const vimFacet = cm.state.facet(Vim) as { mode?: string };
  const currentVimMode = vimFacet?.mode || "unknown";
  console.log("Current Vim Mode:", currentVimMode);
}
