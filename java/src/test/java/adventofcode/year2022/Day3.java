package adventofcode.year2022;

import adventofcode.BaseTest;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

public class Day3 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day3/input.txt");
		int sum = 0;
		for (final String input : data) {
			sum += getSum(input);
		}
		System.out.println(sum);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day3/input.txt");
		int sum = 0;
		for (int i = 0; i < data.size(); i = i + 3) {
			sum += getSumGold(data.get(i), data.get(i + 1), data.get(i + 2));
		}
		System.out.println(sum);
	}

	public int getSumGold(String part1, String part2, String part3) {

		Set<Character> ch1 = part1.chars().mapToObj(e -> (char) e).collect(Collectors.toSet());
		Set<Character> ch2 = part2.chars().mapToObj(e -> (char) e).collect(Collectors.toSet());
		Set<Character> ch3 = part3.chars().mapToObj(e -> (char) e).collect(Collectors.toSet());
		ch1.retainAll(ch2);
		ch1.retainAll(ch3);
		List<Character> item = new ArrayList<>(ch1);
		Character result = item.get(0);
		String resStr = "" + result;
		if (resStr.equals(resStr.toLowerCase())) {
			return result - 'a' + 1;
		}
		return result - 'A' + 27;

	}

	public int getSum(String input) {
		String part1 = input.substring(0, input.length() / 2);
		String part2 = input.substring(input.length() / 2);
		Set<Character> ch1 = part1.chars().mapToObj(e -> (char) e).collect(Collectors.toSet());
		Set<Character> ch2 = part2.chars().mapToObj(e -> (char) e).collect(Collectors.toSet());
		ch1.retainAll(ch2);
		List<Character> item = new ArrayList<>(ch1);
		Character result = item.get(0);
		String resStr = "" + result;
		if (resStr.equals(resStr.toLowerCase())) {
			return result - 'a' + 1;
		}
		return result - 'A' + 27;

	}

}
