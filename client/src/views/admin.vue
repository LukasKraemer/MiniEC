<template>
  USER LOGGED IN

  <button @click="logout">Log out</button>
</template>

<script>
import { TokenStorage } from "@/common/TokenStorage";

export default {
  name: "admin",
  methods: {
    logout: function () {
      TokenStorage.removeToken();
      window.location.href = "/";
    },
  },
};
</script>

<style scoped></style>
