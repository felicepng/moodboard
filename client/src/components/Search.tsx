import type { Component } from 'solid-js';

const Search: Component = () => {
  return (
    <div class="flex justify-center items-center gap-x-4 md:gap-x-10 w-3/4 md:w-1/2 h-8 md:h-10 text-xl">
      <input
        type="text"
        placeholder="Enter theme..."
        class="w-full bg-white text-slate-600 shadow-md shadow-slate-300 h-full focus:outline-none rounded-xl px-5 "
      />
      <button class="bg-white hover:bg-slate-50 text-slate-600 shadow-md shadow-slate-300 h-full rounded-xl px-5">
        search
      </button>
    </div>
  );
};

export default Search;
