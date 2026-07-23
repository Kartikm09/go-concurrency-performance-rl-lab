# tasks/task-001/task_def.py

from typing import List, Dict, Any

class TaskDefinition:
    """Defines the required inputs and outputs for the 'Task 001' problem."""

    @staticmethod
    def get_description() -> str:
        return "Write a function that correctly identifies palindromic sequences in a given list of strings."

    @staticmethod
    def get_signature():
        """Defines the required function signature for candidate solutions."""
        # Input: List[str] (the input data)
        # Output: int (the count of valid palindromes found)
        return "solve(data: List[str]) -> int"

    @staticmethod
    def get_baseline_test_case() -> Dict[str, Any]:
        """Defines the basic inputs and expected outputs."""
        return {
            "input": ["racecar", "hello", "level", "madam"],
            "expected_output": 3  # 'racecar', 'level', 'madam' are palindromes
        }