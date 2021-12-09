package adventofcode.year2021;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;
import java.util.Set;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day9 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day9/input.txt");
		final int xSize = data.size();
		final int ySize = data.get(0).length();

		final int[][] points = new int[xSize][ySize];
		for (int i = 0; i < ySize; i++) {
			for (int j = 0; j < xSize; j++) {
				points[j][i] = Integer.parseInt(data.get(j).charAt(i) + "");
			}
		}
		int count = 0;
		final ArrayList<Integer> resultList = new ArrayList<>();
		for (int i = 0; i < xSize; i++) {
			for (int j = 0; j < ySize; j++) {
				final int current = points[i][j];
				boolean min = true;
				for (int l = -1; l <= 1; l++) {
					for (int m = -1; m <= 1; m++) {
						try {
							if (points[i + l][j + m] < current) {
								min = false;
							}
						} catch (final Exception e) {
						}
					}
				}
				if (min) {
					count = count + current + 1;
				}

			}
		}
		System.out.println(count);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day9/input.txt");
		final int xSize = data.size();
		final int ySize = data.get(0).length();

		final int[][] points = new int[xSize][ySize];
		for (int i = 0; i < ySize; i++) {
			for (int j = 0; j < xSize; j++) {
				points[j][i] = Integer.parseInt(data.get(j).charAt(i) + "");
			}
		}
		final ArrayList<Integer> resultList = new ArrayList<>();
		for (int i = 0; i < xSize; i++) {
			for (int j = 0; j < ySize; j++) {
				final int current = points[i][j];
				boolean min = true;
				for (int l = -1; l <= 1; l++) {
					for (int m = -1; m <= 1; m++) {
						try {
							if (points[i + l][j + m] < current) {
								min = false;
							}
						} catch (final Exception e) {
						}
					}
				}
				if (min) {
					final Set<String> currentBasians = getBasians(new HashSet<>(), points, i, j);
					resultList.add(currentBasians.size());
				}

			}
		}
		Collections.sort(resultList, Collections.reverseOrder());
		System.out.println(resultList.get(0) * resultList.get(1) * resultList.get(2));
	}

	public Set<String> getBasians(final Set<String> basians, final int[][] points, final int i, final int j) {
		final String key = "i" + i + "j" + j + "|" + points[i][j];
		if (basians.contains(key)) {
			return basians;
		} else {
			basians.add(key);
			final int current = points[i][j];
			for (int l = -1; l <= 1; l++) {
				for (int m = -1; m <= 1; m++) {
					try {
						final int check = points[i + l][j + m];
						if (l == 0 || m == 0) {
							if (check < 9 && check > current) {
								final String keyCheck = "i" + (i + l) + "j" + (j + m) + "|" + points[i + l][j + m];
								if (basians.contains(keyCheck)) {
									//already added
								} else {
									basians.addAll(getBasians(basians, points, i + l, j + m));
									basians.add(keyCheck);

								}
							}
						}
					} catch (final Exception e) {
					}
				}
			}
		}
		return basians;
	}

}
