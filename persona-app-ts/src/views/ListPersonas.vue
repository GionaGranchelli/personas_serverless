<template>
  <div class="all-personas">
    <h2>List of Persons</h2>
    <div class="loading" v-if="isLoading">Loading...</div>
    <img alt="Vue logo" class="refresh" src="@/assets/refresh.svg" width="45" height="45" @click="fetchPersons"/>
    <div class="wrapper" v-if="!isLoading">
      <div class="table">
        <div class="row header">
          <div class="cell">ID</div>
          <div class="cell">First Name</div>
          <div class="cell">Last Name</div>
          <div class="cell">Phone Number</div>
          <div class="cell">Address</div>
        </div>
        <div class="row" v-for="person in persons" :key="person.id">
          <div class="cell" data-title="ID">{{ person.id }}</div>
          <div class="cell" data-title="First Name">{{ person.first_name }}</div>
          <div class="cell" data-title="LastN ame">{{ person.last_name }}</div>
          <div class="cell" data-title="Phone Number">{{ person.phone_number }}</div>
          <div class="cell" data-title="Address">{{ person.address }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, inject } from 'vue'
import axios from 'axios';

interface Person {
  id: string;
  first_name: string;
  last_name: string;
  phone_number: string;
  address: string;
}

export default defineComponent({
  name: 'ListPersonas',
  setup() {
    const apiEndpoint = inject<string>('apiEndpoint');
    const persons = ref<Person[]>([]);
    const isLoading = ref<boolean>(false);
    const fetchPersons = async () => {
      try {
        isLoading.value = true;
        const response = await axios.get(`${apiEndpoint}/personas`);
        persons.value = response.data;
        isLoading.value = false;
      } catch (error) {
        isLoading.value = false;
        console.log('Failed to retrieve persons.');
        console.error(error);
        alert('Failed to retrieve persons.');
        persons.value = [];
      }
    };

    onMounted(fetchPersons);

    return {
      persons,
      isLoading,
      fetchPersons
    };
  },
});
</script>
<style scoped>
.all-personas {
  margin-top: 3rem;
  min-height: 100vh;
  box-sizing: border-box;
  background: rgba(0,0,0,.5);
  box-shadow: 0 15px 25px rgba(0,0,0,.6);
  border-radius: 10px;
}
.all-personas h2 {
  margin: 0 0 30px;
  padding-top: 40px;
  color: #fff;
  text-align: center;
}
.refresh {
  display: block;
  cursor: pointer;
  margin: 0 auto 1%;
}
.wrapper .loading {
  margin: 0 auto;
  max-width: 800px;
  border-radius: 10px;
  text-align: center;
}

.table {
  margin: 0 0 40px 0;
  width: 100%;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
  display: table;
  border-radius: 10px;
}
@media screen and (max-width: 300px) {
  .table {
    display: block;
    border-radius: 10px;
  }
}

.row {
  display: table-row;
  background: #f6f6f6;
}
.row:nth-of-type(even):hover {
  background: #636161;
}
.row:nth-of-type(odd):hover:not(.header) {
  background: #636161;
}
.row:nth-of-type(odd) {
  background: #e9e9e9;
}
.row.header {
  font-weight: 900;
  color: #ffffff;
  background: hsla(160, 100%, 37%, 1);
}
.row.green {
  background: hsla(160, 100%, 37%, 1);
}
.row.blue {
  background: #2980b9;
}
@media screen and (max-width: 580px) {
  .row {
    padding: 14px 0 7px;
    display: block;
  }
  .row.header {
    padding: 0;
    height: 6px;
  }
  .row.header .cell {
    display: none;
  }
  .row .cell {
    margin-bottom: 10px;
  }
  .row .cell:before {
    margin-bottom: 3px;
    content: attr(data-title);
    min-width: 98px;
    font-size: 10px;
    line-height: 10px;
    font-weight: bold;
    text-transform: uppercase;
    color: #969696;
    display: block;
  }
}

.cell {
  padding: 6px 12px;
  display: table-cell;
  color: black;
}
@media screen and (max-width: 580px) {
  .cell {
    padding: 2px 16px;
    display: block;
  }
}
</style>