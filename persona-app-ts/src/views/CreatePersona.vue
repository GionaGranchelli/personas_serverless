<template>
  <div class="persona-box">
    <h2>Create Person</h2>
    <form @submit.prevent="createPerson">
      <div class="control">
        <input type="text" v-model="person.first_name" required/>
        <label for="first_name">First Name:</label>
      </div>
      <div class="control">
        <input type="text" v-model="person.last_name" required />
        <label for="last_name">Last Name:</label>
      </div>
      <div class="control">
        <input type="text" v-model="person.phone_number" required />
        <label for="phone_number">Phone Number:</label>
      </div>
      <div class="control">
        <input type="text" v-model="person.address" required />
        <label for="address">Address:</label>
      </div>
     <div>
       &nbsp;
     </div>
      <button type="submit">Create</button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, inject, reactive } from 'vue'
import axios from 'axios'

interface Person {
  id?: string;
  first_name: string;
  last_name: string;
  phone_number: string;
  address: string;
}

export default defineComponent({
  name: 'CreatePersona',
  setup() {
    const person = reactive<Person>({
      first_name: '',
      last_name: '',
      phone_number: '',
      address: ''
    })
    const apiEndpoint = inject<string>('apiEndpoint')
    const createPerson = async () => {
      try {
        console.log(person)
        const response = await axios.post(`${apiEndpoint}personas`, person)
        console.log(response.data)
        alert('Person created successfully!')
        Object.assign(person, { id: '', first_name: '', last_name: '', phone_number: '', address: '' })
      } catch (error) {
        console.error(error)
        alert('Failed to create person.')
      }
    }

    return {
      person,
      createPerson,
      apiEndpoint
    }
  }
})
</script>
<style>
.persona-box {
  padding: 40px;
  background: rgba(0,0,0,.5);
  box-sizing: border-box;
  box-shadow: 0 15px 25px rgba(0,0,0,.6);
  border-radius: 10px;
  margin-top: 3rem;
  min-height: 100vh;
}

.persona-box h2 {
  margin: 0 0 30px;
  padding: 0;
  color: #fff;
  text-align: center;
}

.persona-box .control {
  position: relative;
}

.persona-box .control input {
  width: 100%;
  padding: 10px 0;
  font-size: 16px;
  color: #fff;
  margin-bottom: 30px;
  border: none;
  border-bottom: 1px solid #fff;
  outline: none;
  background: transparent;
}
.persona-box .control label {
  position: absolute;
  top:0;
  left: 0;
  padding: 10px 0;
  font-size: 16px;
  color: #fff;
  pointer-events: none;
  transition: .5s;
}

.persona-box .control input:focus ~ label,
.persona-box .control input:valid ~ label {
  top: -20px;
  left: 0;
  color: hsla(160, 100%, 37%, 1);
  font-size: 12px;
}

.persona-box form a {
  position: relative;
  display: inline-block;
  padding: 10px 20px;
  color: hsla(160, 100%, 37%, 1);
  font-size: 16px;
  text-decoration: none;
  text-transform: uppercase;
  overflow: hidden;
  transition: .5s;
  margin-top: 40px;
  letter-spacing: 4px
}

</style>