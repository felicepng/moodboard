import type { Component } from 'solid-js';
import { setTheme } from '../App';

const Search: Component = () => {
  const onSearch = (event: Event) => {
    event.preventDefault();
    const input = (event.target as HTMLFormElement)['0'].value;
    setTheme([input || 'random theme']);
  };

  return (
    <form
      onSubmit={onSearch}
      class="flex justify-center items-center gap-x-4 md:gap-x-10 w-3/4 md:w-1/2 h-8 md:h-10 text-xl md:text-2xl"
    >
      <input
        type="text"
        placeholder="enter a theme..."
        class="w-full bg-white text-slate-600 shadow-md shadow-slate-300 h-full focus:outline-none rounded-xl px-5 "
      />
      <button
        type="submit"
        class="bg-white hover:bg-slate-50 text-slate-600 shadow-md shadow-slate-300 h-full rounded-xl px-5"
      >
        search
      </button>
    </form>
  );
};

export default Search;
