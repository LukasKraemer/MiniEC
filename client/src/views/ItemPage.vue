<template>
  <img
    v-bind:alt="name"
    v-bind:src="picture"
    style="width: 500px; height: 500px"
  />
  <br />
  Name: {{ name }} <br />
  Price: {{ price }} €<br />
  ID: {{ id }} <br />
  <router-link to="/"><button>Back</button></router-link>
</template>

<script>
import axios from "axios";
import { Global } from "@/common/global";

export default {
  name: "ItemPage",
  created: function () {
    this.load();
  },
  data() {
    return {
      name: "ARTIKELNAME",
      price: "-1" + "€",
      idfromroute: this.$route.params.id,
      id: -1,
      picture: "",
    };
  },
  methods: {
    load: function () {
      axios
        .get(Global.APIURL + "/product/" + this.idfromroute)
        .then((response) => {
          console.log(response.data);
          this.name = response.data[0].name;
          this.price = response.data[0].price;
          this.id = response.data[0].id;
          this.picture = Global.APIURL + "img/" + response.data[0].picture;
        })
        .catch((error) => {
          this.product = error;
        });
    },
  },
};
</script>

<style scoped></style>
