package adventofcode.year2019;

import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day18 extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2019/day18/input_sample.txt");
		int count = 0;
		for (final String input : data) {
			count = input.length();
			System.out.println(input);
		}
		System.out.println(count);
	}

}

