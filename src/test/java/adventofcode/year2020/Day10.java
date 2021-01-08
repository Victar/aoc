package adventofcode.year2020;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day10 extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<Integer> data = readIntFromFile("year2020/day10/input.txt");
		int countX1 = 1;
		int countX2 = 0;
		int countX3 = 0;
		Collections.sort(data);
		System.out.println(data);

		for (int i = 0; i < data.size(); i++) {
			if (i == data.size() - 1) {
				countX3 = countX3 + 1;
			} else {
				int diff = data.get(i + 1) - data.get(i);
				if (diff == 1) {
					countX1 = countX1 + 1;

				} else if (diff == 3) {
					countX3 = countX3 + 1;

				} else {
					System.out.println("not valid");
				}
			}
		}
		//		Collections.sort(data, new Comparator<Integer>() {})
		System.out.println(countX1 + " " + countX2 + " " + countX3);
		System.out.println(countX1 * countX3);

	}

	@Test public void runGold() throws Exception {
		final ArrayList<Integer> data = readIntFromFile("year2020/day10/input.txt");
		data.add(0);
		Collections.sort(data);
		data.add(data.get(data.size() - 1) + 3);
		System.out.println(data);
		ArrayList<Long> pathList = new ArrayList<>();
		for (int i = 0; i < data.size(); i++) {
			pathList.add((long) -1);
		}
		System.out.println(pathList);
		System.out.println(countPaths(data, pathList, 0));
		System.out.println(pathList);


	}

	public long countPaths(ArrayList<Integer> data, ArrayList<Long> pathList, int i) {
		if (i == data.size()-1) {
			return 1;
		} else if (pathList.get(i) > -1) {
			return pathList.get(i);
		}
		long pathesForI = 0;
		for (int j = i+1; j < Math.min(data.size(), j + 3); j++) {
			if (data.get(j) - data.get(i) <= 3) {
//				System.out.println(i + " - " + j + "  " + data.get(i) + " " + data.get(j));
				pathesForI += countPaths(data, pathList, j);
			}
		}
		pathList.set(i, pathesForI);
		return pathesForI;
	}
}

