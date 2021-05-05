<template>
  <h1>Lukas Online Shop</h1>
  <div id="nav">
    <router-link to="/">Shop</router-link> |
    <router-link to="/login">Login</router-link>
  </div>
  <router-view />
  <div class="spacer"></div>
  <MyFooter></MyFooter>
</template>

<style scoped>
h1 {
  margin-top: 1em;
}
</style>

<style>
.spacer {
  flex: 1;
}
body {
  margin: 0;
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
<script>
import MyFooter from "@/components/MyFooter";
export default {
  components: { MyFooter },
};
</script>
