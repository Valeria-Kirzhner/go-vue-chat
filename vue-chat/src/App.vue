<template>
  <div class="container">
    <div class="d-flex flex-column align-items-stretch flex-shrink-0 bg-white">
      <div
        class="d-flex align-items-center flex-shrink-0 p-3 link-dark text-decoration-none border-bottom"
      >
        <input class="fs-5 fw-semibold" v-model="username" />
      </div>
      <div class="list-group list-group-flush border-bottom scrollarea">
        <a
          class="list-group-item list-group-item-action py-3 lh-sm"
          v-for="message in messages"
          :key="message"
        >
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">{{ message.username }}</strong>
            <small class="text-muted">{{ currentDate }}</small>
          </div>
          <div class="col-10 mb-1 small">{{ message.message }}</div>
        </a>
      </div>
    </div>
    <form @submit.prevent="onSubmit">
      <input
        class="form-controll"
        placeholder="Write a message"
        v-model="message"
      />
    </form>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";
import Pusher from "pusher-js";

export default {
  name: "App",
  setup() {
    const username = ref("username");
    const messages = ref([]);
    const message = ref("");
    const currentDate = ref("");

    onMounted(() => {
      Pusher.logToConsole = true;

      const pusher = new Pusher("88553fa8f6ec1411c857", {
        cluster: "ap2",
      });

      const channel = pusher.subscribe("chat");
      channel.bind("message", (data) => {
        getDate();
        messages.value.push(data);
      });
    });

    function getDate() {
      const current = new Date();
      const time = `${current.getHours()}:${current.getMinutes()}:${current.getSeconds()}`;
      const dateAndTime = `${time}  ${current.getDate()}/${
        current.getMonth() + 1
      }/${current.getFullYear()}`;
      return (currentDate.value = dateAndTime);
    }

    const onSubmit = async () => {
      await fetch("http://localhost:8000/api/messages", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          username: username.value,
          message: message.value,
        }),
      });
      message.value = "";
    };

    return {
      username,
      messages,
      message,
      onSubmit,
      currentDate,
    };
  },
};
</script>

<style>
.scrollarea {
  min-height: 500px;
}
</style>
