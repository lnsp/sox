<template>
  <div class="h-full flex flex-col">
    <div class="flex max-w-xl justify-between items-center">
      <headline h1>Machine</headline>
    </div>
    <divider light />
    <div class="bg-white border-2 border-gray-300 rounded-lg flex-grow flex flex-col">
      <div class="border-b-2 border-gray-300 p-4 flex">
        Hello
      </div>
      <div class="flex flex-grow">
        <div class="w-48 border-r-2 border-gray-300">
          <div class="p-4">
            <machine-nav-item to="networking">Networking</machine-nav-item>
            <machine-nav-item to="storage">Storage</machine-nav-item>
            <machine-nav-item to="power">Power</machine-nav-item>
          </div>
        </div>
        <div class="flex-grow">

        </div>
      </div>

    </div>
  </div>
</template>