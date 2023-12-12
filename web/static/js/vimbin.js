// Desc: Main JavaScript file for the Vimbin web app
document.addEventListener("DOMContentLoaded", function () {
  // Function to get the preferred theme (dark, light, or system default)
  function getPreferredTheme() {
    // If the browser doesn't support the prefers-color-scheme media query, return the default theme
    if (!window.matchMedia) {
      return "default";
    }

    if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
      return "catppuccin";
    }

    if (window.matchMedia("(prefers-color-scheme: light)").matches) {
      return "default";
    }

    return "default";
  }

  // Function to update Vim mode display
  function updateVimMode(vimEvent, vimModeElement) {
    const mode = vimEvent.mode;
    const sub = vimEvent.subMode;

    vimModeElement.classList.remove(
      "normal",
      "insert",
      "visual",
      "visual-line",
    );

    switch (mode) {
      case "normal":
        vimModeElement.innerText = "NORMAL";
        vimModeElement.classList.add("normal");
        break;
      case "insert":
        vimModeElement.innerText = "INSERT";
        vimModeElement.classList.add("insert");
        break;
      case "visual":
        if (sub === "") {
          vimModeElement.innerText = "VISUAL";
          vimModeElement.classList.add("visual");
          break;
        }
        vimModeElement.innerText = "V-LINE";
        vimModeElement.classList.add("visual-line");
        break;
      default:
        vimModeElement.innerText = "UNKNOWN";
        vimModeElement.classList.add("unknown");
    }
  }

  var editor = CodeMirror.fromTextArea(document.getElementById("code"), {
    lineNumbers: true,
    mode: "text/x-csrc",
    keyMap: "vim",
    matchBrackets: true,
    showCursorWhenSelecting: true,
    theme: getPreferredTheme(),
  });

  editor.focus();

  var vimMode = document.getElementById("vim-mode");
  CodeMirror.on(editor, "vim-mode-change", function (e) {
    updateVimMode(e, vimMode);
  });

  CodeMirror.commands.save = async function () {
    let status = "No changes were made.";

    try {
      const response = await fetch("/save", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ content: editor.getValue() }),
      });

      if (!response.ok) {
        throw new Error("Save failed. Reason: " + response.statusText);
      }

      // Check if the response has a valid JSON body
      const isJson = response.headers
        .get("content-type")
        ?.includes("application/json");

      if (!isJson) {
        throw new Error("Response was not JSON");
      }

      const changesResponse = await response.json();

      if (changesResponse.status !== "no changes") {
        // Retrieve the number of bytes written from the response headers
        const bytesWritten = response.headers.get("X-Bytes-Written");
        status = `${bytesWritten}B written`;
      }
    } catch (error) {
      status = "Error saving: " + error.message;
    }
    document.getElementById("vim-command-line").innerText = status;
  };

  // Listen for changes in the prefers-color-scheme media query
  window.matchMedia("(prefers-color-scheme: dark)").addListener((e) => {
    if (e.matches) {
      editor.setOption("theme", "catppuccin");
    } else {
      editor.setOption("theme", "default");
    }
  });

  window.matchMedia("(prefers-color-scheme: light)").addListener((e) => {
    if (e.matches) {
      editor.setOption("theme", "default");
    }
  });
});
