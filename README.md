# Trivial Concurrency Exercises
http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/

### Exercise 1: The Daily Walk

Every morning, Alice and Bob go for a walk, and being creatures of habit, they follow the same routine every day.

First, they both prepare, grabbing sunglasses, perhaps a belt, closing open windows, turning off ceiling fans, and pocketing their phones and keys.

Once they’re both ready, which typically takes each of them between 60 and 90 seconds, they arm the alarm, which has a 60 second delay.

While the alarm is counting down, they both put on their shoes, a process which tends to take each of them between 35 and 45 seconds.

Then they leave the house together and lock the door, before the alarm has finished its countdown.

*Write a program to simulate Alice and Bob’s morning routine.*

Here’s some sample output from running a solution to this problem.

```
Let's go for a walk!
Bob started getting ready
Alice started getting ready
Alice spent 72 seconds getting ready
Bob spent 76 seconds getting ready
Arming alarm.
Bob started putting on shoes
Alarm is counting down.
Alice started putting on shoes
Alice spent 37 seconds putting on shoes
Bob spent 39 seconds putting on shoes
Exiting and locking the door.
Alarm is armed.

```

### Exercise 2: Eating Tapas
Four friends are having dinner together. They ordered five dishes to share, each of which consists of between 5 and 10 morsels.

They eat leisurely, spending between 30 seconds and 3 minutes eating each morsel.

*Write a program to simulate the meal.*

Running the program might look something like this:

```
Bon appétit!
Alice is enjoying some chorizo
Bob is enjoying some chopitos
Charlie is enjoying some pimientos de padrón
Dave is enjoying some croquetas
Alice is enjoying some patatas bravas
Charlie is enjoying some chorizo
Dave is enjoying some chopitos
Alice is enjoying some pimientos de padrón
Bob is enjoying some croquetas
Dave is enjoying some patatas bravas
Alice is enjoying some chorizo
Bob is enjoying some chopitos
Charlie is enjoying some pimientos de padrón
Alice is enjoying some croquetas
Bob is enjoying some patatas bravas
Dave is enjoying some chorizo
Charlie is enjoying some chopitos
Alice is enjoying some pimientos de padrón
Dave is enjoying some patatas bravas
Charlie is enjoying some croquetas
Bob is enjoying some chorizo
Alice is enjoying some chopitos
Charlie is enjoying some pimientos de padrón
Charlie is enjoying some patatas bravas
Dave is enjoying some croquetas
Alice is enjoying some chorizo
Bob is enjoying some chopitos
Bob is enjoying some pimientos de padrón
Charlie is enjoying some patatas bravas
Alice is enjoying some chopitos
Dave is enjoying some patatas bravas
Charlie is enjoying some chopitos
Alice is enjoying some patatas bravas
Bob is enjoying some patatas bravas
Dave is enjoying some patatas bravas
That was delicious!
```
### Exercise 3: Internet Café
A small internet café in a village just outside of Manilla has 8 computers, which are available on a first-come first-serve basis. When all the computers are taken, the next person in line has to wait until a computer frees up.

This morning several groups of tourists, 25 people in all, are waiting when the doors open.

Each person spends between 15 minutes and 2 hours online.

*Write a program to simulate the computer usage at the café*

Sample output:
```
Tourist 4 is online.
Tourist 25 is online.
Tourist 9 is online.
Tourist 15 is online.
Tourist 1 is online.
Tourist 16 is online.
Tourist 2 is online.
Tourist 5 is online.
Tourist 17 waiting for turn.
Tourist 6 waiting for turn.
Tourist 3 waiting for turn.
Tourist 18 waiting for turn.
Tourist 7 waiting for turn.
Tourist 12 waiting for turn.
Tourist 13 waiting for turn.
Tourist 10 waiting for turn.
Tourist 19 waiting for turn.
Tourist 14 waiting for turn.
Tourist 11 waiting for turn.
Tourist 20 waiting for turn.
Tourist 22 waiting for turn.
Tourist 23 waiting for turn.
Tourist 24 waiting for turn.
Tourist 8 waiting for turn.
Tourist 21 waiting for turn.
Tourist 2 is done, having spent 15 minutes online.
Tourist 17 is online.
Tourist 5 is done, having spent 18 minutes online.
Tourist 6 is online.
Tourist 15 is done, having spent 40 minutes online.
Tourist 3 is online.
Tourist 16 is done, having spent 41 minutes online.
Tourist 18 is online.
Tourist 6 is done, having spent 40 minutes online.
Tourist 7 is online.
Tourist 4 is done, having spent 60 minutes online.
Tourist 12 is online.
Tourist 17 is done, having spent 64 minutes online.
Tourist 13 is online.
Tourist 1 is done, having spent 88 minutes online.
Tourist 10 is online.
Tourist 9 is done, having spent 100 minutes online.
Tourist 19 is online.
Tourist 25 is done, having spent 114 minutes online.
Tourist 14 is online.
Tourist 19 is done, having spent 24 minutes online.
Tourist 11 is online.
Tourist 7 is done, having spent 88 minutes online.
Tourist 20 is online.
Tourist 18 is done, having spent 107 minutes online.
Tourist 22 is online.
Tourist 3 is done, having spent 110 minutes online.
Tourist 23 is online.
Tourist 12 is done, having spent 100 minutes online.
Tourist 24 is online.
Tourist 13 is done, having spent 87 minutes online.
Tourist 8 is online.
Tourist 20 is done, having spent 26 minutes online.
Tourist 21 is online.
Tourist 11 is done, having spent 65 minutes online.
Tourist 24 is done, having spent 32 minutes online.
Tourist 8 is done, having spent 28 minutes online.
Tourist 10 is done, having spent 110 minutes online.
Tourist 23 is done, having spent 57 minutes online.
Tourist 14 is done, having spent 98 minutes online.
Tourist 22 is done, having spent 106 minutes online.
Tourist 21 is done, having spent 111 minutes online.
The place is empty, let's close up and go to the beach!
```
### Notes

Use *time.Sleep()* to simulate delays and activities.

Remember to seed math/rand so that people’s activities (eating morsels and putting on shoes) take different amounts of time each time you run the code.

The exercises were inspired by William Kennedy’s post **The Nature Of Channels In Go.**




