<script lang="ts">
  import Button from "./lib/Button.svelte";
  import QuizCard from "./lib/QuizCard.svelte";
  ////////////////////////////////////////
  let quizzes: { _id: string; name: string }[] = [];
  async function getQuizzes() {
    let response = await fetch("http://localhost:5000/api/quizzes");
    if (!response.ok) {
      return "Failed to fetch quizzes";
    }

    let jsonData = await response.json();
    quizzes = jsonData;
  }

  function hostQuiz(quiz) {
    console.log(quiz);
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
  {#each quizzes as quiz}
    <QuizCard on:host={() => hostQuiz(quiz)} {quiz} />
  {/each}
  <div></div>
  <Button>HII</Button>
</main>

<style>
</style>
