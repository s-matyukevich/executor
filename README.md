Simple task excutor.

Main features:
1. Allows to execute tasks sequentially or in parallel.
2. Handle task status (Waiting, Running, Failed, Expired, Finished)
3. Allws to concurently track status of executing tasksk
4. Applys timout strategy for each task and to whole execution

Sample ussage:

	stage1 := NewStage(
		[]*Task{
			&Task{
				Name: "task1-1",
				Func: func(){
					fmt.Print("Task 1-1 is executing\n")
				},
			},
			&Task{
				Name: "task1-2",
				Func: func(){
					fmt.Print("Task 1-2 is executing\n")
				},
			},
			&Task{
				Name: "task1-3",
				Func: func(){
					fmt.Print("Task 1-3 is executing\n")
				},
			},
		}, false)
	stage2 := NewStage(
		[]*Task{
			&Task{
				Name: "task2-1",
				Func: func(){
					fmt.Print("Task 2-1 is executing\n")
				},
			},
			&Task{
				Name: "task2-2",
				Func: func(){
					fmt.Print("Task 2-2 is executing\n")
				},
			},
			&Task{
				Name: "task2-3",
				Func: func(){
					fmt.Print("Task 2-3 is executing\n")
				},
			},
		}, true)
	executor := &Executor{}
	executor.AddStage(stage1)
	executor.AddStage(stage2)
	executor.Execute()
