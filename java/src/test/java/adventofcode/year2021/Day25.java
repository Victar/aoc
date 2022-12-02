package adventofcode.year2021;

import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day25 extends BaseTest {

	boolean canMove = true;
	int SIZE_X;
	int SIZE_Y;

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day25/input.txt");
		this.SIZE_X = data.get(0).length();
		this.SIZE_Y = data.size();
		Cucumber[][] cucumbers = new Cucumber[this.SIZE_X][this.SIZE_Y];
		for (int y = 0; y < this.SIZE_Y; y++) {
			for (int x = 0; x < this.SIZE_X; x++) {
				final String currentS = "" + data.get(y).charAt(x);
				if (!".".equals(currentS)) {
					cucumbers[x][y] = new Cucumber(currentS, x, y);
				}
			}
		}
		int roundCount = 0;
		//		printArray(cucumbers);
		while (this.canMove) {
			roundCount++;
			cucumbers = move(cucumbers);
		}
		printArray(cucumbers);

		System.out.println(roundCount);
	}

	public Cucumber[][] move(final Cucumber[][] cucumbers) {
		final Cucumber[][] cucumbersNew = new Cucumber[this.SIZE_X][this.SIZE_Y];
		boolean changed = false;
		for (int y = 0; y < this.SIZE_Y; y++) {
			for (int x = 0; x < this.SIZE_X; x++) {
				final Cucumber current = cucumbers[x][y];
				if (current != null) {
					if (current.east && cucumbers[(x + 1) % this.SIZE_X][y] == null) {
						cucumbersNew[(x + 1) % this.SIZE_X][y] = current;
						changed = true;
					} else {
						cucumbersNew[x][y] = current;
					}
				}
			}
		}
		final Cucumber[][] cucumbersNew2 = new Cucumber[this.SIZE_X][this.SIZE_Y];
		for (int y = 0; y < this.SIZE_Y; y++) {
			for (int x = 0; x < this.SIZE_X; x++) {
				final Cucumber current = cucumbersNew[x][y];
				if (current != null) {
					if (!current.east && cucumbersNew[x][(y + 1) % this.SIZE_Y] == null) {
						cucumbersNew2[x][(y + 1) % this.SIZE_Y] = current;
						changed = true;
					} else {
						cucumbersNew2[x][y] = current;
					}
				}
			}
		}
		if (!changed) {
			this.canMove = false;
		}
		return cucumbersNew2;
	}

	public void printArray(final Cucumber[][] array) {

		System.out.println("----------");
		for (int y = 0; y < this.SIZE_Y; y++) {
			for (int x = 0; x < this.SIZE_X; x++) {
				final Cucumber current = array[x][y];
				System.out.print(current == null ? "." : current.east ? ">" : "v");
			}
			System.out.print(" y ---" + y);
			System.out.println();
		}
	}

	@Data class Cucumber {

		boolean east;
		int x;
		int y;

		public Cucumber(final String type, final int x, final int y) {
			this.east = ">".equals(type);
			this.x = x;
			this.y = y;
		}
	}

}
