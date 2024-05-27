# The Stacks - A Library Simulator

> **_Finally, a chance to live my dream of developing library management software!_**
> — John Bannister (and possibly no-one else ever)

The Stacks is (or maybe I should say 'will be') a simulation of a council library which has recently been furnished with a new library management system. Your job is to develop that library management system!

Librarians and service users will use various devices (computers, self-serve checkout machines, sensor gates etc.) to stock, move, borrow and return books. These devices will make API requests to your <abbr title="Library Management System">LMS</abbr> and the responses will influence the resulting actions taken by the people in the simulation.

## Constraints

To keep the project focused, I will be working with the following constraints:

- **Language** - [Go]

- **Framework** - [raylib] with [raylib-go] bindings

- **Geometry** - All game logic will act on a 2D plane (physics, navigation, interaction etc.)

- **UI** - [raygui] with [raylib-go/raygui] bindings

[Go]: https://go.dev
[raygui]: https://github.com/raysan5/raygui
[raylib-go]: https://github.com/gen2brain/raylib-go
[raylib-go/raygui]: https://github.com/gen2brain/raylib-go/tree/master/raygui
[raylib]: https://www.raylib.com
