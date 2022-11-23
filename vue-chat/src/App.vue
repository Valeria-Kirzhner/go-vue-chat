<template>
  <div class="container">
    <div class="d-flex flex-column align-items-stretch flex-shrink-0 bg-white">
      <a
        href="/"
        class="d-flex align-items-center flex-shrink-0 p-3 link-dark text-decoration-none border-bottom"
      >
        <input class="fs-5 fw-semibold" v-model="username" />
      </a>
      <div class="list-group list-group-flush border-bottom scrollarea">
        <a
          class="list-group-item list-group-item-action py-3 lh-sm"
          v-for="message in messages"
          :key="message"
        >
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">{{ message.username }}</strong>
            <small class="text-muted"></small>
          </div>
          <div class="col-10 mb-1 small">
            Some placeholder content in a paragraph below the heading and date.
          </div>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
export default {
  name: "App",
  setup() {
    const username = ref("username");
    const messages = ref([]);

    return {
      username,
      messages,
    };
  },
};
</script>

<style></style>
