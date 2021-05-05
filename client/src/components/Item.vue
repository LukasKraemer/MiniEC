<template>

    <div class="col-xs-6 col-sm-4 col-md-3" style="margin-bottom: 5%">
        <div style="border: 3px solid grey">
      <img
        v-bind:alt="name"
        v-bind:src="picture"
        style="
          width: 100%;
          height: auto;
        "
      />
      <hr>
      <div style="background: #42b983; bottom: 0">
      <span style="left: 0; text-align: left"> Name:{{ itemName }}</span>
      <span style="right: 0; text-align: right">Price: {{ price }}€</span>
      </div>
            </div>
    </div>

</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";

@Options({
  props: {
    price: Number,
    itemName: String,
    picture: String,
  },
})
export default class Item extends Vue {
  price!: number;
  itemName!: string;
  picture!: string;
}
</script>

<style scoped>
</style>
