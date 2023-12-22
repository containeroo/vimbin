import { initializeEditor } from "./editor-setup";
import "./vim-commands";

declare var apiToken: string;

async function getContent() {
  try {
    const response = await fetch("/fetch", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "X-API-Token": apiToken,
      },
    });

    if (!response.ok) {
      throw new Error(`Fetch failed. Reason: ${response.statusText}`);
    }
    const content = await response.text(); // Read the response body as text
    console.log("Raw Response:", content);

    return content;
  } catch (error) {
    const statusElement = document.getElementById("status");
    if (!statusElement) {
      console.error("Status element not found");
      return "";
    }

    clearTimeout((statusElement as any).timerId); // Clear the existing timer before setting a new one

    statusElement.innerText = `ERROR: ${(error as Error).message}`;
    statusElement.classList.add("isError");

    // Set a new timer
    (statusElement as any).timerId = setTimeout(() => {
      statusElement.innerText = "";
      statusElement.classList.remove("isError");
    }, 5000);
    return "";
  }
}

async function main() {
  await initializeEditor(await getContent());
}

main();
