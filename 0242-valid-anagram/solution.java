class Solution {
    public boolean isAnagram(String s, String t) {
        HashMap<Character, Integer> m = new HashMap<>();

        for (char x : s.toCharArray()) {
            m.put(x, m.getOrDefault(x, 0) + 1);
        }

        for (char x : t.toCharArray()) {
            m.put(x, m.getOrDefault(x, 0) - 1);
        }

        for (int val : m.values()) {
            if (val != 0) {
                return false;
            }
        }

        return true;
    }
}
