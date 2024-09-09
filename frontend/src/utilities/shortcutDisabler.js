export const shortcutDisabler = () => {
  document.addEventListener("contextmenu", (event) => event.preventDefault());
  document.addEventListener("keydown", (event) => {
    if (
      event.key === "F5" ||
      (event.key.toLowerCase() === "r" && event.ctrlKey) ||
      (event.key.toLowerCase() === "f" && event.ctrlKey)
    ) {
      event.preventDefault();
    }
  });
};