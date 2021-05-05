<template>
  <div style="border: 3px solid darkgrey" class="row bg-light bg-gradient">
    <div class="col-md-4 text-muted">
      <label for="type">What Kind of</label><br />
      <select
        class="select text-secondary"
        id="type"
        v-model="cat"
        @change="updateForm"
      >
        <option value="1">Computer</option>
        <option value="2">Laptop</option>
        <option value="3">Server</option>
      </select>
    </div>
    <div class="col-md-4 text-muted">
      <Slider
        @change="updateForm"
        :format="format"
        :step="10"
        :min="sliderRange[0]"
        :max="sliderRange[1]"
        v-model="value"
      />
    </div>
    <div class="col-md-4 text-success">
      SEARCH <br /><input
        type="search"
        name="search"
        v-model="search"
        @keydown="updateForm"
      />
    </div>
  </div>
  <div class="container" @load="this.updateForm">
    <div class="row items">
      <template v-for="item in product" :key="item.id">
        <Item
          @click="loadElement($event, item.id)"
          :id="item.id"
          :itemName="item.name"
          :price="item.price"
          :picture="item.picture"
        ></Item>
        <br />
      </template>
    </div>
  </div>
</template>

<script>
import Item from "@/components/Item.vue";
import axios from "axios";
import { Global } from "@/common/global";
import Slider from "@vueform/slider";

export default {
  components: {
    Item,
    Slider,
  },
  data() {
    return {
      product: [],
      value: [100, 2000],
      search: "",
      cat: 1,
      sliderRange: [0,10000],
      format: {
        suffix: " €",
        decimals: 2,
      },
    };
  },
  created() {
    this.updateForm();
  },
  methods: {
    loadElement: function (e, id) {
      this.$router.push({ name: "article", params: { id: id } });
    },
    updateForm: function () {
      axios
        .get(Global.APIURL + "product/", {
          params: {
            minPrice: this.value[0],
            maxPrice: this.value[1],
            name: this.search,
            category: this.cat,
          },
        })
        .then((response) => {
          this.product = response.data;

          this.product.forEach((element) => {
            element.picture = Global.APIURL + "img/" + element.picture;
          });
        })
        .catch((error) => {
          this.product = error;
        });
    },
  },
};
</script>

<style src="@vueform/slider/themes/default.css"></style>

<style>
.slider-base,
.slider-connects {
  margin-top: 5%;
}
</style>
<style scoped>
.row .items {
  height: auto;
}

form {
  display: grid;
  align-items: center;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
}

.items {
  padding-top: 3%;
}
</style>
