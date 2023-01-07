import type { Component } from 'solid-js';
import { For } from 'solid-js';
import { images } from '../App';

const Prompts: Component = () => {
  return (
    <div class="flex flex-col gap-y-4 items-center text-center text-lg md:text-2xl text-slate-400">
      <p>generated image prompts:</p>
      <div class="grid md:grid-cols-2 md:gap-x-20">
        <For each={images().prompts}>
          {(prompt: string, idx) => (
            <p>
              {idx() + 1}. {prompt}
            </p>
          )}
        </For>
      </div>
    </div>
  );
};

export default Prompts;
