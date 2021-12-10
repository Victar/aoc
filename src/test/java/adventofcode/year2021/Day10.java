package adventofcode.year2021;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Stack;
import java.util.stream.Collectors;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day10 extends BaseTest {

	private static HashMap<Character, Integer> createMap() {
		final HashMap<Character, Integer> result = new HashMap<>();
		result.put(')', 3);
		result.put(']', 57);
		result.put('}', 1197);
		result.put('>', 25137);
		return result;
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day10/input.txt");
		final HashMap<Character, Integer> dataResult = createMap();
		int result = 0;
		for (final String input : data) {
			final Character c = getCharacter(input);
			if (dataResult.containsKey(c)) {
				result += dataResult.get(c);
			}
		}
		System.out.println(result);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day10/input.txt");
		final List<Long> scores = new ArrayList<>();
		for (final String input : data) {
			if (isValid(input)) {
				scores.add(Long.parseLong(StringUtils.replaceEach(getRestString(input), new String[] { ")", "]", "}", ">" },
						new String[] { "1", "2", "3", "4" }), 5));
			}
		}
		Collections.sort(scores);
		System.out.println(scores.get(scores.size() / 2));
	}

	public String getRestString(final String s) {
		final Stack<Character> stack = new Stack<>();
		for (final char c : s.toCharArray()) {
			if (c == '(') stack.push(')');
			else if (c == '{') stack.push('}');
			else if (c == '<') stack.push('>');
			else if (c == '[') stack.push(']');
			else if (stack.isEmpty() || stack.pop() != c) {
			}
		}
		Collections.reverse(stack);
		return stack.stream().map(c -> c.toString()).collect(Collectors.joining());
	}

	public Character getCharacter(final String s) {
		final Stack<Character> stack = new Stack<>();
		for (final char c : s.toCharArray()) {
			if (c == '(') stack.push(')');
			else if (c == '{') stack.push('}');
			else if (c == '<') stack.push('>');
			else if (c == '[') stack.push(']');
			else if (stack.isEmpty() || stack.pop() != c) return c;
		}

		return null;
	}

	public boolean isValid(final String s) {
		final Stack<Character> stack = new Stack<>();
		for (final char c : s.toCharArray()) {
			if (c == '(') stack.push(')');
			else if (c == '{') stack.push('}');
			else if (c == '<') stack.push('>');
			else if (c == '[') stack.push(']');
			else if (stack.isEmpty() || stack.pop() != c) return false;
		}
		return true;
	}

}
