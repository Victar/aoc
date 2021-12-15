package adventofcode.year2021;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import org.junit.Test;

import adventofcode.BaseTest;
import adventofcode.GraphUtil;
import lombok.Data;

public class Day15 extends BaseTest {

	@Test public void runSilver() throws Exception {
		runAny(1);
	}

	@Test public void runGold() throws Exception {
		runAny(5);
	}

	public void runAny(final int size) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day15/input.txt");

		final int SIZE_Y = data.size();
		final int SIZE_X = data.get(0).length();
		final int SIZE_Y_G = SIZE_Y * size;
		final int SIZE_X_G = SIZE_X * size;

		final Node[][] NODES = new Node[SIZE_X_G][SIZE_Y_G];

		for (int y = 0; y < SIZE_Y_G; y++) {
			for (int x = 0; x < SIZE_X_G; x++) {
				final int value = (Integer.parseInt("" + data.get(x % SIZE_X).charAt(y % SIZE_Y)) + (x / SIZE_X) + (y / SIZE_Y)) % 9;
				final int risk = value == 0 ? 9 : value;
				final Node currentNode = new Node(x, y, risk);
				NODES[x][y] = currentNode;
			}
		}

		//set neighbors
		for (int y = 0; y < SIZE_Y_G; y++) {
			for (int x = 0; x < SIZE_X_G; x++) {
				for (int l = x - 1; l <= x + 1; l++) {
					for (int m = y - 1; m <= y + 1; m++) {
						if (l == x || m == y) {
							if (l >= 0 && l < SIZE_X_G && m >= 0 && m < SIZE_Y_G && (l != x || m != y)) {
								NODES[x][y].addDestination(NODES[l][m]);
							}
						}
					}
				}
			}
		}
		final Node source = NODES[0][0];
		GraphUtil.calculatePath(source, 0);
		System.out.println(NODES[SIZE_X_G - 1][SIZE_Y_G - 1].getDistance());
		printArray(NODES, SIZE_X_G, SIZE_Y_G, NODES[SIZE_X_G - 1][SIZE_Y_G - 1].getMinPath());
	}

	public void printArray(final Node[][] array, final int SIZE_X_G, final int SIZE_Y_G, final List<GraphUtil.Node> path) {
		for (int x = 0; x < SIZE_X_G; x++) {
			for (int y = 0; y < SIZE_Y_G; y++) {
				final boolean inPath = path.contains(new Node(x, y, 0));
				if (inPath) {
					System.out.print("(" + array[x][y].getRisk() + ")");
				} else {
					System.out.print(" " + array[x][y].getRisk() + " ");
				}
			}
			System.out.println();
		}
	}

	@Data public class Node implements GraphUtil.Node {

		int x;
		int y;
		int risk;
		private List<GraphUtil.Node> minPath = new LinkedList<>();
		private Integer distance = Integer.MAX_VALUE;

		Map<GraphUtil.Node, Integer> neighbors = new HashMap<>();

		public Node(final int x, final int y, final int risk) {
			this.x = x;
			this.y = y;
			this.risk = risk;
		}

		@Override public int getDistance() {
			return this.distance;
		}

		@Override public void setDistance(final int distance) {
			this.distance = distance;
		}

		public void addDestination(final Node destination) {
			this.neighbors.put(destination, destination.getRisk());
		}

		@Override public boolean equals(final Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;

			final Node node = (Node) o;

			if (this.x != node.x) return false;
			return this.y == node.y;
		}

		@Override public int hashCode() {
			int result = this.x;
			result = 31 * result + this.y;
			return result;
		}

		@Override public void setMinPath(final List<GraphUtil.Node> minPath) {
			this.minPath = minPath;
		}

		@Override public Map<GraphUtil.Node, Integer> getNeighbors() {
			return this.neighbors;
		}
	}
}
