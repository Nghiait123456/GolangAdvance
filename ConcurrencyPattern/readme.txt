/*
 Concurrency pattern in golang from golang development team, other resource, other frame... and my experience
 Concurrency is simple run and not simple stop. It has many problem with code concurrency.
 It's not dependence anything frame or packet.
 It is maybe use in everything in golang code.
 It's use for me as an as doc research and search.
 I didn't create it, it's not clear who first created it, but I synthesized it and deepened it for my work.
 I will it's hope help someone up to up golang skills, i'm very happy if it happened.
 Enjoy and relax it.
 My contact: minhnghia.pham.it@gmail.com
*/

/*
  Basic concurrency understand before learn ConcurrencyPattern's part:
  1) What Concurrency?
  2) Why use concurrency?
  3) Distinctive concurrency and parallelism?
  4) What is problem concurrency then it's not strong?

  ==================================================================================================================================================
  1) What Concurrency?:
     Concurrency is the composition of independently executing computations.
     Concurrency is a way to structure software, particularly as a way to write clean code that interacts well with the real world.
     In Wikipedia: "concurrency is the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or in partial order, without affecting the final outcome."

     ==> Summary: There are some problem, we can it executed on many part, maybe is many core or not, ==> summary target : problem run faster but still correct.

    ex: Http WebServer: in one server, we handle many request incoming. To achieve highest for performance, every webserver have solution for concurrency difference.
    in Golang, summary solution is use routine run worker handle request, there are 3 main way implement:
               1) every routine for every request.
               2) pool routine for all request ( in at time, one routine handle one request)
               3) pool routine with event(Epool, ...)

  2) Why use concurrency?
     Concurrency promoting strength in problems: high load, repeating, not much logic interwoven between tasks in one problem.
     This problems are many times in present: handle request, handle socket, handle message queue, call api, handle big data ...
     This is reason concurrency is increasingly popular.
     In Golang, concurrency implement in routine, one routine cost is 2 to 8 kb ==> mini cost for concurrency.


  3) Distinctive concurrency and parallelism?
     For wiki: "Task parallelism (also known as function parallelism and control parallelism) is a form of parallelization of computer code across multiple processors in parallel computing environments."
     This explanation is not simple for understand it and understand difference with concurrency.
     There are simple way, follow this case:
        Assume problem run on computer has only one core CPU.
        You are never way implement parallelism in this context.
        But you have way implement concurrency in this context.
     ===>    Think about this context, you will understand Distinctive concurrency and parallelism.


  4) What is problem concurrency then it's not strong?
      If bottleneck of problems is not dependency concurrency, is dependency for other condition (io, disk,...), have constraint condition in this task concurrency.
          Ex: you update inventory, you implement one lock for every update. When high load, you can up number concurrency but over range limit of mutex.
          Have many routine race on mutex, but performance is not greater.
          routine 1, routine 2, ... routine n wait and race to have mutex.
          Time and cost for switch context and routine 1, routine 2... is large.
          In this context, concurrency is slower simple thread run sync( and not use mutex)

      ==> In this problems, children task run on routine is not run independence, problems is not really for concurrency, it is need other solution.
   ======================================================================================================================================================================================================
*/
