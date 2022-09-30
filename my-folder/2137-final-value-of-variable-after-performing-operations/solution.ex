defmodule Solution do
  @spec final_value_after_operations(operations :: [String.t]) :: integer
  def final_value_after_operations(operations) do
    Enum.reduce(operations, 0, fn x, acc -> acc + get_operator(x) end)
  end
  
  defp get_operator(operator) when operator in ["X++", "++X"], do: 1
  defp get_operator(operator) when operator in ["X--", "--X"], do: -1
  defp get_operator(_), do: 0
end
