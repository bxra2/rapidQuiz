<script lang="ts">
  import svelteLogo from "./assets/svelte.svg";
  import viteLogo from "/vite.svg";
  import Counter from "./lib/Counter.svelte";
  import Button from "./lib/Button.svelte";
  ////////////////////////////////////////
  async function getQuizzes() {
    let response = await fetch("http://localhost:5000/api/quizzes");
    if (!response.ok) {
      return "Failed to fetch quizzes";
    }

    let jsonData = await response.json();

    console.log(jsonData);
  }

  function connect() {
    let websocket = new WebSocket("ws://localhost:5000/ws");
    websocket.onopen = () => {
      console.log("opened connection");
      websocket.send("hello peeps");
    };

    websocket.onmessage = (event) => {
      console.log(event.data);
    };
  }
</script>

<main>
  <button on:click={getQuizzes}>get Quizzes</button>
  <button on:click={connect}>connect</button>

  <div></div>
  <Button />
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
    color: #888;
  }
</style>
