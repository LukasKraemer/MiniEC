<template>
  <form @submit.prevent="login">
    <div class="box">
      <h1>Login</h1>
      <input
        class="mail"
        placeholder="E-Mail Adress"
        type="email"
        name="email"
        id="email"
        v-model="email"
      />
      <br />
      <input
        placeholder="Password"
        class="mail"
        type="password"
        name="passwd"
        id="passwd"
        v-model="passwd"
      />
      <br />
      <input class="btn" type="submit" value="LOGIN" />
    </div>
  </form>
</template>

<script>
import axios from "axios";
import { Global } from "@/common/global";
import { TokenStorage } from "@/common/TokenStorage";
import { useRouter, useRoute } from "vue-router";

export default {
  name: "Login",
  created: function () {
    if (TokenStorage.getToken() != null) {
      window.location.href = "/admin";
    }
  },
  data() {
    return {
      email: "",
      passwd: "",
      token: TokenStorage.getToken(),
      router: useRouter(),
      route: useRoute(),
    };
  },
  methods: {
    login: function () {
      axios
        .get(Global.APIURL + "login", {
          headers: {
            Authorization: "Basic " + btoa(this.email + ":" + this.passwd),
          },
        })
        .then(function (response) {
          TokenStorage.setToken(response.data);
          window.location.href = "/admin";
        })
        .catch(function (error) {
          console.log(error);
        });
    },
  },
};
</script>

<style scoped>
p {
  font-size: 12px;
  text-decoration: none;
  color: #ffffff;
}

h1 {
  margin-top: 15px;
  font-size: 1.5em;
  color: #525252;
}

.box {
  background: white;
  width: 300px;
  border-radius: 6px;
  margin: 0 auto 0 auto;
  padding: 0 0 30px 0;
  border: #2980b9 4px solid;
}

.email {
  background: #ecf0f1;
  border: #ccc 1px solid;
  border-bottom: #ccc 2px solid;
  padding: 8px;
  width: 250px;
  color: #aaaaaa;
  margin-top: 10px;
  font-size: 1em;
  border-radius: 4px;
}

.password {
  border-radius: 4px;
  background: #ecf0f1;
  border: #ccc 1px solid;
  padding: 8px;
  width: 250px;
  font-size: 1em;
}

.btn {
  background: #2ecc71;
  width: 60%;
  padding-top: 5px;
  padding-bottom: 5px;
  color: white;
  border-radius: 4px;
  border: #27ae60 1px solid;
  margin-top: 20px;
  margin-bottom: 20px;
  font-weight: 800;
  font-size: 0.8em;
}

.btn:hover {
  background: #2cc06b;
}

#btn2 {
  float: left;
  background: #3498db;
  width: 125px;
  padding-top: 5px;
  padding-bottom: 5px;
  color: white;
  border-radius: 4px;
  border: #2980b9 1px solid;

  margin-top: 20px;
  margin-bottom: 20px;
  margin-left: 10px;
  font-weight: 800;
  font-size: 0.8em;
}

form {
  margin-bottom: 25%;
}
</style>
