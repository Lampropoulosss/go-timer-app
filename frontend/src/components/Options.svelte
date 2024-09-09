<script>
  // @ts-nocheck

  import Buttons from "./Buttons.svelte";

  // @ts-nocheck
  export let timerStarted = false;
  export let isPaused = false;
  export let alarmFinished = false;

  let timer = "";
  let isInvalid = false;

  function handleKeyDown(e) {
    if (
      e.target.value.length > 8 ||
      !/^[0-9]$/i.test(e.target.value[e.target.value.length - 1])
    ) {
      e.target.value = e.target.value.slice(0, e.target.value.length - 1);
    } else if (!e.key) {
      if (e.target.value.length == 2) {
        timer = e.target.value + ":";
      } else if (e.target.value.length == 5) {
        timer = e.target.value + ":";
      } else {
        timer = e.target.value;
      }
    } else if (e.key === "Backspace") {
      if (e.target.value.length == 6) {
        timer = e.target.value.slice(0, 5);
      } else if (e.target.value.length == 3) {
        timer = e.target.value.slice(0, 2);
      } else {
        timer = e.target.value.slice(0, e.target.value.length);
      }
    }
  }
</script>

<div>
  <input
    type="text"
    placeholder="HH:MM:SS"
    value={timer}
    on:input={(e) => {
      isInvalid = false;
      handleKeyDown(e);
    }}
    on:keydown={(e) => {
      if (e.key === "Backspace") handleKeyDown(e);
    }}
    class:is-invalid={isInvalid}
    autocomplete="off"
    disabled={timerStarted}
  />
  <Buttons
    {timer}
    {isInvalid}
    {alarmFinished}
    {isPaused}
    {timerStarted}
    on:clear
    on:pause-resume
    on:start
    on:invalid={() => {
      isInvalid = true;
    }}
  />
</div>

<style>
  ::placeholder {
    text-align: center;
  }

  div {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.5em;
  }

  input {
    border-radius: 8px;
    border: 2px solid transparent;
    padding: 0.4 0.5em;
    font-family: Pacifico;
    font-size: 1.1em;
    font-weight: 500;
    color: #ffffff;
    background-color: #0f0f0f98;
    transition: border-color 0.25s;
    box-shadow: 0 2px 2px rgba(0, 0, 0, 0.2);
    width: 100px;
    text-align: center;
    outline: none;
  }

  input::placeholder {
    font-size: 0.8em;
    font-family: cursive;
  }

  input:hover {
    border-color: #00ced1;
  }

  input:disabled:hover {
    border-color: transparent;
    pointer-events: none;
  }

  input.is-invalid {
    border-color: rgba(220, 20, 60, 0.658);
  }
</style>
