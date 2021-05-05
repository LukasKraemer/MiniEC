<template>
  <footer class="bg-light text-center text-lg-start">
    <div class="text-center">
      © 2021 Copyright:
      <a class="text-dark" href="https://lukas-kraemer.de/">Lukas Krämer</a>
    </div>
    <div class="text-center">
      <a class="text-dark" v-bind:href="imprint">Impressum</a>
    </div>
    <div class="text-center">
      <a class="text-dark" v-bind:href="dataPolicy">Datenschutz</a>
    </div>
  </footer>
</template>

<script>
import { Global } from "@/common/global";
export default {
  name: "Footer",
  data: function () {
    return {
      imprint: Global.Imprint,
      dataPolicy: Global.dataPolicy,
    };
  },
};
</script>

<style scoped>
.text-center {
  background-color: rgba(0, 0, 0, 0.2);
}
</style>
