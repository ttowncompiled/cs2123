(time go run hw2.go books/huckleberryFinn.txt bin_results/huckleberryFinn) > bin_results/huckleberryFinn.time 2>&1 && \
(time go run hw2.go books/tomSawyer.txt bin_results/tomSawyer) > bin_results/tomSawyer.time 2>&1 && \
(time go run hw2.go books/innocentAdventuress.txt bin_results/innocentAdventuress) > bin_results/innocentAdventuress.time 2>&1 && \
(time go run hw2.go books/muchAdoAboutNothing.txt bin_results/muchAdoAboutNothing) > bin_results/muchAdoAboutNothing.time 2>&1 && \
(time go run hw2.go books/prodigalVillage.txt bin_results/prodigalVillage) > bin_results/prodigalVillage.time 2>&1 && \
(time go run hw2-avl.go books/huckleberryFinn.txt avl_results/huckleberryFinn) > avl_results/huckleberryFinn.time 2>&1 && \
(time go run hw2-avl.go books/tomSawyer.txt avl_results/tomSawyer) > avl_results/tomSawyer.time 2>&1 && \
(time go run hw2-avl.go books/innocentAdventuress.txt avl_results/innocentAdventuress) > avl_results/innocentAdventuress.time 2>&1 && \
(time go run hw2-avl.go books/muchAdoAboutNothing.txt avl_results/muchAdoAboutNothing) > avl_results/muchAdoAboutNothing.time 2>&1 && \
(time go run hw2-avl.go books/prodigalVillage.txt avl_results/prodigalVillage) > avl_results/prodigalVillage.time 2>&1
