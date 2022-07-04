// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

contract Todo {
    address public owner;

    struct Task {
        string content;
        bool completed;
    }

    Task[] tasks;

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function addTask(string memory _content) public onlyOwner {
        tasks.push(Task(_content, false));
    }

    function getTask(uint256 _id) public view onlyOwner returns (Task memory) {
        return tasks[_id];
    }

    function listTasks() public view onlyOwner returns (Task[] memory) {
        return tasks;
    }

    function updateTask(uint256 _id, string memory _content) public onlyOwner {
        tasks[_id].content = _content;
    }

    function toggleTaskCompletion(uint256 _id) public onlyOwner {
        tasks[_id].completed = !tasks[_id].completed;
    }

    function deleteTask(uint256 _id) public onlyOwner {
        // left-shift tasks from index _id
        for (uint256 i = _id; i < tasks.length - 1; i++) {
            tasks[_id] = tasks[_id + 1];
        }

        // then pop last task as it is now a duplicate
        tasks.pop();
    }
}
