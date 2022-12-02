package adventofcode.year2021;

import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day1 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day1/input.txt");
		int count = -1;
		int current = 0;
		int previous = 0;

		for (final String password : data) {
			current = Integer.parseInt(password);
			if (current > previous) {
				count++;
			}
			previous = current;
		}
		System.out.println(count);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day1/input.txt");
		int count = -1;
		int current1, current2, current3 = 0;
		final ArrayList<Integer> data3 = new ArrayList<Integer>();
		for (int i = 0; i < data.size(); i++) {
			current1 = Integer.parseInt(data.get(i));
			if (i < data.size() - 1) {
				current2 = Integer.parseInt(data.get(i + 1));
			} else {
				current2 = 0;
			}
			if (i < data.size() - 2) {
				current3 = Integer.parseInt(data.get(i + 2));
			} else {
				current3 = 0;
			}
			data3.add(current1 + current2 + current3);
		}
		int current = 0;
		int previous = 0;

		for (final Integer password : data3) {
			current = password;
			if (current > previous) {
				count++;
			}
			previous = current;
		}
		System.out.println(count);
	}

}
