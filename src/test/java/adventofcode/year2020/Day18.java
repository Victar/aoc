package adventofcode.year2020;

import java.util.*;

import org.apache.commons.lang3.StringUtils;
import org.junit.Ignore;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day18 extends BaseTest {

	private static final Map<String, Integer> prioritySilver;
	private static final Map<String, Integer> priorityGold;

	static {
		final Map<String, Integer> temp = new HashMap<>();
		temp.put("*", 1);
		temp.put("+", 1);
		temp.put("(", 0);
		prioritySilver = Collections.unmodifiableMap(temp);
	}

	static {
		final Map<String, Integer> temp = new HashMap<>();
		temp.put("*", 1);
		temp.put("+", 2);
		temp.put("(", 0);
		priorityGold = Collections.unmodifiableMap(temp);
	}

	@Test @Ignore public void singleCheck() {
		//		System.out.println(calcString("1 + 2 + 3", prioritySilver));

		System.out.println(calcString("1 + (2 * 3) + (4 * (5 + 6))", prioritySilver));
		//System.out.println(calcString("5 + (8 + 3 * 9 + 3 * 4 * 3)", prioritySilver));
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day18/input.txt");
		long count = 0;
		for (final String input : data) {
			count += calcString(input, prioritySilver);
		}
		System.out.println(count);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day18/input.txt");
		long count = 0;
		for (final String input : data) {
			count += calcString(input, priorityGold);
		}
		System.out.println(count);
	}

	public long calcString(String input, final Map<String, Integer> priorityMap) {
		input = input.replaceAll("\\(", "\\( ").replaceAll("\\)", " \\)");
		//		System.out.println("Infix: " + input);
		final String[] notation = StringUtils.split(input, " ");
		//		System.out.println("Infix array: " + Arrays.toString(notation));
		final Queue<String> queue = convertInfixToRPN(notation, priorityMap);
		//		System.out.println("RPN array: " + queue);
		return evalRPN(queue.toArray(new String[] {}));

	}

	boolean isNumber(final String str) {
		try {
			Long.valueOf(str);
			return true;
		} catch (final Exception e) {
			return false;
		}
	}

	public Queue<String> convertInfixToRPN(final String[] infixNotation, final Map<String, Integer> priority) {
		final Queue<String> queue = new LinkedList<>();
		final Stack<String> stack = new Stack<>();
		for (String token : infixNotation) {
			token = token.trim();
			//			System.out.println("token:" + token +  " stack:" + stack + " queue:" +queue);
			if ("(".equals(token)) {
				stack.push(token);
				continue;
			}

			if (")".equals(token)) {
				while (!"(".equals(stack.peek())) {
					queue.add(stack.pop());
				}
				stack.pop();
				continue;
			}
			if (priority.containsKey(token)) {
				while (!stack.empty() && priority.get(token) <= priority.get(stack.peek())) {
					queue.add(stack.pop());
				}
				stack.push(token);
				continue;
			}
			if (isNumber(token)) {
				queue.add(token);
				continue;
			}
			throw new IllegalArgumentException("Invalid input");
		}
		while (!stack.isEmpty()) {
			queue.add(stack.pop());
		}
		return queue;
	}

	public long evalRPN(final String[] tokens) {
		long returnValue = 0;
		final String operators = "+*";
		final Stack<String> stack = new Stack<String>();
		for (final String token : tokens) {
			if (!operators.contains(token)) {
				stack.push(token);
			} else {
				final long a = Long.valueOf(stack.pop());
				final long b = Long.valueOf(stack.pop());
				switch (token) {
					case "+":
						stack.push(String.valueOf(a + b));
						break;
					case "*":
						stack.push(String.valueOf(a * b));
						break;

				}
			}
			//			System.out.println("RPN: "  + stack);
		}

		returnValue = Long.valueOf(stack.pop());

		return returnValue;
	}

}

