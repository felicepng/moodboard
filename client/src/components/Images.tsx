import type { Component } from 'solid-js';
import { For } from 'solid-js';

const __urls = new Array(8).fill('');

const Images: Component = () => {
  return (
    <div class="grid grid-cols-2 md:grid-cols-4 gap-8">
      <For each={__urls}>
        {(url) => (
          <div class="bg-gray-500 w-32 h-32 md:w-40 md:h-40 2xl:w-60 2xl:h-60" />
        )}
      </For>
    </div>
  );
};

export default Images;
