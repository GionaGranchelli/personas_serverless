<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Welcome from '@/components/Welcome.vue'

</script>

<template>
  <header>
    <div class="wrapper">
      <Welcome msg="You did it!" />
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
        <RouterLink to="/create-persona">Create Persona</RouterLink>
        <RouterLink to="/list-personas">List Persona</RouterLink>
        <RouterLink to="/search-persona">Search Persona</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
header {
  display: grid;
  grid-auto-columns: max-content;
  padding-right: calc(var(--section-gap) / 2);
  grid-auto-rows: auto;
}
header .wrapper {
  display: grid;
  grid-auto-flow: row;
}

nav {
  text-align: left;
  margin-left: -1rem;
  font-size: 1rem;
  padding: 1rem 0;
  margin-top: 1rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}


</style>
