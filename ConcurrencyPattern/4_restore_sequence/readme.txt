/*
ideal:
 it'sBad same job in 3_generator, but we want when we need data, we will get it, and not, routine is pending.
 we have flag share read and write data.

    imagine, we have pool save flag:

          |       |
          |       |
          |       |
          |       |
          | flat3 |
          | flat2 |
          |_flat1_|

    every time after routine write data stream, we need one data from flag channel to continue (flag1, flag2, flag3, ...), if not is pending
    every time after get data from data stream, we push one data to flag channel, it'sBad fuel (flag1, flag2, flag3, ...) for all routine stream continue.
    receive and write to flag is consecutive, if only one pending, routine is pending
*/
