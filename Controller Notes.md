#Writing a Controller Notes

Watch events → enqueue work → reconcile item → talk to API server to create/update/delete resources → repeat.

Common building blocks: informer/watch, workqueue (rate-limiting), client (k8s API), and the reconcile loop. Controller tooling: client-go (low level) or controller-runtime/Kubebuilder (higher level).

Two ways to implement 

a) client-go (low level) — you wire informers, listers, workqueue yourself. Teaches internals, more code.

b) controller-runtime / Kubebuilder (high level) — provides Manager, cached client, controller scaffolding, and reconciler interface. Faster to build production controllers; recommended unless you need toy internals

=====================

4 percent []3-2-1 

2 percent 2-1 

0.75 percent []

=========
5.875 
600k 
30 years 
weehawken 

3549.23 
[buydown] 3-2-1 17393 [2490 2822 3176 3549 ]
2-1 8814 []
1-1 5984 []

5.875- 5.75 1.023 of 6.1k

5.875 - 4.99 [4.24*6=254]