package adventofcode.year2019;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day5 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2019/day2/input.txt");
		final List<Integer> result = new ArrayList<Integer>();
		final String[] arr = data.get(0).split(",");
		for (int i = 0; i < arr.length; i++) {
			result.add(Integer.parseInt(arr[i]));
		}
		result.set(1, 12);
		result.set(2, 2);
		boolean process = true;
		for (int i = 0; i < result.size() && process; i = i + 4) {
			final int current = result.get(i);
			if (current == 99) {
				process = false;
			} else if (current == 1) {
				result.set(result.get(i + 3), result.get(result.get(i + 1)) + result.get(result.get(i + 2)));
			} else if (current == 2) {
				result.set(result.get(i + 3), result.get(result.get(i + 1)) * result.get(result.get(i + 2)));
			}
		}
		System.out.println(result.get(0));
	}


}
