[English](README.md) · **简体中文**

# pipeline

先 fan-out 各阶段，再 fan-in 合并。每个阶段为 `func(in <-chan T) <-chan U`。

Tags: `concurrency` `fan-in` `fan-out`
