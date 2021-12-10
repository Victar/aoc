package adventofcode.year2021;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Stack;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day10 extends BaseTest {

	private static HashMap<String, Integer> createMap() {
		final HashMap<String, Integer> result = new HashMap<>();
		result.put(")", 3);
		result.put("]", 57);
		result.put("}", 1197);
		result.put(">", 25137);
		return result;
	}

	private static HashMap<String, Integer> createMapG() {
		final HashMap<String, Integer> result = new HashMap<>();
		result.put(")", 1);
		result.put("]", 2);
		result.put("}", 3);
		result.put(">", 4);
		return result;
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day10/input.txt");
		final HashMap<String, Integer> dataResult = createMap();
		int result = 0;
		for (final String input : data) {
			final String validStr = isValidStr(input);
			if (dataResult.containsKey(validStr)) {
				result += dataResult.get(validStr);
			}
		}
		System.out.println(result);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day10/input.txt");
		final HashMap<String, Integer> dataResult = createMapG();
		final List<Long> scores = new ArrayList<>();
		for (final String input : data) {
			long current = 0;
			final boolean isValid = isValid(input);
			if (isValid) {
				final List<Character> stack = validStrG(input);
				if (stack.size() > 0) {
					for (int i = stack.size() - 1; i != -1; i--) {
						final String currentS = "" + stack.get(i);
						final int cur = dataResult.get(currentS);
						current = current * 5 + cur;
					}
					scores.add(current);
				}
			}
		}
		Collections.sort(scores);
		System.out.println(scores.get(scores.size() / 2));
	}

	public ArrayList<Character> validStrG(final String s) {
		final Stack<Character> stack = new Stack<>();
		final int size = 0;
		for (final char c : s.toCharArray()) {
			if (c == '(') stack.push(')');
			else if (c == '{') stack.push('}');
			else if (c == '<') stack.push('>');
			else if (c == '[') stack.push(']');
			else if (stack.isEmpty() || stack.pop() != c) {
			}
		}
		return new ArrayList<>(stack);
	}

	public String isValidStr(final String s) {
		final Stack<Character> stack = new Stack<>();
		for (final char c : s.toCharArray()) {
			if (c == '(') stack.push(')');
			else if (c == '{') stack.push('}');
			else if (c == '<') stack.push('>');
			else if (c == '[') stack.push(']');
			else if (stack.isEmpty() || stack.pop() != c) return c + "";
		}

		return StringUtils.EMPTY;
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
